package ai.verta.modeldb.entities;

public interface ComponentEntity {

  void setBlobHash(String elementSha);

  void setBaseBlobHash(String folderHash);
}
