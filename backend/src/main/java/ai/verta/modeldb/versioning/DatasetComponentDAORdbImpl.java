package ai.verta.modeldb.versioning;

import ai.verta.modeldb.entities.ComponentEntity;
import ai.verta.modeldb.entities.dataset.PathDatasetComponentBlobEntity;
import ai.verta.modeldb.entities.dataset.S3DatasetComponentBlobEntity;
import ai.verta.modeldb.entities.versioning.InternalFolderElementEntity;
import java.security.NoSuchAlgorithmException;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.UUID;
import org.hibernate.Session;

public class DatasetComponentDAORdbImpl implements DatasetComponentDAO {

  public static final String TREE = "TREE";

  static class TreeElem {
    String path;
    String sha256 = null;
    String type = null;
    Map<String, TreeElem> children = new HashMap<>();

    TreeElem() {}

    TreeElem push(List<String> pathList, String sha256, String type) {
      path = pathList.get(0);
      if (pathList.size() > 1) {
        children.putIfAbsent(pathList.get(1), new TreeElem());
        this.type = TREE;
        return children
            .get(pathList.get(1))
            .push(pathList.subList(1, pathList.size()), sha256, type);
      } else {
        this.sha256 = sha256;
        this.type = type;
        return this;
      }
    }

    String getPath() {
      return path != null ? path : "";
    }

    String getSha256() {
      return sha256;
    }

    String getType() {
      return type;
    }

    InternalFolderElement saveFolders(Session session, FileHasher fileHasher)
        throws NoSuchAlgorithmException {
      if (children.isEmpty()) {
        return InternalFolderElement.newBuilder()
            .setElementName(getPath())
            .setElementSha(getSha256())
            .build();
      } else {
        InternalFolder.Builder internalFolder = InternalFolder.newBuilder();
        List<InternalFolderElement> elems = new LinkedList<>();
        for (TreeElem elem : children.values()) {
          InternalFolderElement build = elem.saveFolders(session, fileHasher);
          elems.add(build);
          if (elem.getType().equals(TREE)) {
            internalFolder.addSubFolders(build);
          } else {
            internalFolder.addBlobs(build);
          }
        }
        final InternalFolderElement build =
            InternalFolderElement.newBuilder()
                .setElementName(getPath())
                .setElementSha(fileHasher.getSha(internalFolder.build()))
                .build();
        Iterator<TreeElem> iter = children.values().iterator();
        for (InternalFolderElement elem : elems) {
          session.saveOrUpdate(
              new InternalFolderElementEntity(elem, build.getElementSha(), iter.next().getType()));
        }
        return build;
      }
    }
  }

  @Override
  public String setBlobs(Session session, List<BlobExpanded> blobsList, FileHasher fileHasher)
      throws NoSuchAlgorithmException {
    TreeElem treeElem = new TreeElem();
    List<ComponentEntity> componentEntities = new LinkedList<>();
    for (BlobExpanded blob : blobsList) {
      final DatasetBlob dataset = blob.getBlob().getDataset();
      final String[] split = blob.getPath().split("/");
      TreeElem treeChild = treeElem.push(Arrays.asList(split), fileHasher.getSha(dataset), TREE);
      switch (dataset.getContentCase()) {
        case S3:
          for (S3DatasetComponentBlob componentBlob : dataset.getS3().getComponentsList()) {
            final String sha256 = componentBlob.getPath().getSha256();
            S3DatasetComponentBlobEntity s3DatasetComponentBlobEntity =
                new S3DatasetComponentBlobEntity(
                    UUID.randomUUID().toString(), sha256, componentBlob.getPath());
            componentEntities.add(s3DatasetComponentBlobEntity);
            treeChild.push(
                Arrays.asList(split[split.length - 1], componentBlob.getPath().getPath()),
                componentBlob.getPath().getSha256(),
                componentBlob.getClass().getSimpleName());
          }
          break;
        case PATH:
          for (PathDatasetComponentBlob componentBlob : dataset.getPath().getComponentsList()) {
            final String sha256 = componentBlob.getSha256();
            PathDatasetComponentBlobEntity pathDatasetComponentBlobEntity =
                new PathDatasetComponentBlobEntity(
                    UUID.randomUUID().toString(), sha256, componentBlob);
            componentEntities.add(pathDatasetComponentBlobEntity);
            treeChild.push(
                Arrays.asList(split[split.length - 1], componentBlob.getPath()),
                componentBlob.getSha256(),
                componentBlob.getClass().getSimpleName());
          }
          break;
      }
    }
    final InternalFolderElement internalFolderElement = treeElem.saveFolders(session, fileHasher);
    session.saveOrUpdate(new InternalFolderElementEntity(internalFolderElement, null, TREE));
    final String elementSha = internalFolderElement.getElementSha();
    for (ComponentEntity componentEntity : componentEntities) {
      session.saveOrUpdate(componentEntity);
    }
    return elementSha;
  }
}
