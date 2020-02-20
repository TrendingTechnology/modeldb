package ai.verta.modeldb.versioning;

import java.security.NoSuchAlgorithmException;
import java.util.List;
import org.hibernate.Session;

interface DatasetComponentDAO {

  String setBlobs(Session session, List<BlobExpanded> blobsList, FileHasher fileHasher)
      throws NoSuchAlgorithmException;
}
