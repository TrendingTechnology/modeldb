package ai.verta.modeldb.entities.dataset;

import ai.verta.modeldb.entities.ComponentEntity;
import ai.verta.modeldb.versioning.PathDatasetComponentBlob;
import java.io.Serializable;
import java.util.Objects;
import javax.persistence.Column;
import javax.persistence.Embeddable;
import javax.persistence.EmbeddedId;
import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Table(name = "path_dataset_component_blob")
public class PathDatasetComponentBlobEntity implements ComponentEntity {
  public PathDatasetComponentBlobEntity() {}

  public PathDatasetComponentBlobEntity(
      String blobHash, PathDatasetComponentBlob pathDatasetComponentBlob) {

    this.id = new PathDatasetComponentBlobId(blobHash);
    this.path = pathDatasetComponentBlob.getPath();
    this.size = pathDatasetComponentBlob.getSize();
    this.last_modified_at_source = pathDatasetComponentBlob.getLastModifiedAtSource();
    this.sha256 = pathDatasetComponentBlob.getSha256();
    this.md5 = pathDatasetComponentBlob.getMd5();
  }

  @EmbeddedId private PathDatasetComponentBlobId id;

  @Column(name = "path", columnDefinition = "TEXT")
  private String path;

  @Column(name = "size")
  private Long size;

  @Column(name = "last_modified_at_source")
  private Long last_modified_at_source;

  @Column(name = "sha256", columnDefinition = "text")
  private String sha256;

  @Column(name = "md5", columnDefinition = "text")
  private String md5;

  public String getPath() {
    return path;
  }

  public Long getSize() {
    return size;
  }

  public Long getLast_modified_at_source() {
    return last_modified_at_source;
  }

  public String getSha256() {
    return sha256;
  }

  public String getMd5() {
    return md5;
  }

  public PathDatasetComponentBlob toProto() {
    return PathDatasetComponentBlob.newBuilder()
        .setPath(this.path)
        .setSize(this.size)
        .setLastModifiedAtSource(this.last_modified_at_source)
        .setSha256(this.sha256)
        .setMd5(this.md5)
        .build();
  }

  @Override
  public void setBlobHash(String blobHash) {
    id.setBlob_hash(blobHash);
  }

  @Override
  public void setBaseBlobHash(String folderHash) {
    id.setPath_dataset_blob_id(folderHash);
  }
}

@Embeddable
class PathDatasetComponentBlobId implements Serializable {

  @Column(name = "blob_hash", nullable = false, columnDefinition = "varchar", length = 64)
  private String blob_hash;

  @Column(
      name = "path_dataset_blob_id",
      nullable = false,
      columnDefinition = "varchar",
      length = 64)
  private String path_dataset_blob_id;

  public PathDatasetComponentBlobId(String blobHash) {
    this.blob_hash = blobHash;
  }

  private PathDatasetComponentBlobId() {}

  public String getBlob_hash() {
    return blob_hash;
  }

  public String getPath_dataset_blob_id() {
    return path_dataset_blob_id;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (!(o instanceof S3DatasetComponentBlobId)) return false;
    S3DatasetComponentBlobId that = (S3DatasetComponentBlobId) o;
    return Objects.equals(getBlob_hash(), that.getBlob_hash())
        && Objects.equals(getPath_dataset_blob_id(), that.getS3_dataset_blob_id());
  }

  @Override
  public int hashCode() {
    return Objects.hash(getBlob_hash(), getPath_dataset_blob_id());
  }

  public void setBlob_hash(String blob_hash) {
    this.blob_hash = blob_hash;
  }

  public void setPath_dataset_blob_id(String path_dataset_blob_id) {
    this.path_dataset_blob_id = path_dataset_blob_id;
  }
}
