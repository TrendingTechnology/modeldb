package ai.verta.modeldb.versioning;

import ai.verta.modeldb.ModelDBException;
import java.security.NoSuchAlgorithmException;
import org.hibernate.Session;

public interface BlobFunction {
  String apply() throws NoSuchAlgorithmException, ModelDBException;
}
