package ai.verta.modeldb.entities.hyperparameterConfig;

import ai.verta.modeldb.versioning.HyperparameterValuesConfigBlob;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "hyperparameter_element_config_blob")
public class HyperparameterElementConfigBlobEntity {
  private HyperparameterElementConfigBlobEntity() {}

  public HyperparameterElementConfigBlobEntity(
      String blobHash,
      String commitHash,
      String name,
      Integer valueType,
      HyperparameterValuesConfigBlob hyperparameterConfigBlob) {
    this.blob_hash = blobHash;
    this.name = name;
    this.commit_hash = commitHash;
    this.value_type = valueType;
    this.int_value = hyperparameterConfigBlob.getIntValue();
    this.float_value = hyperparameterConfigBlob.getFloatValue();
    this.string_value = hyperparameterConfigBlob.getStringValue();
  }

  @Id
  @Column(name = "blob_hash", columnDefinition = "varchar", length = 64, nullable = false)
  private String blob_hash;

  @Column(name = "name", columnDefinition = "varchar")
  private String name;

  @Column(name = "commit_hash", columnDefinition = "varchar", length = 64)
  private String commit_hash;

  @Column(name = "value_type")
  private Integer value_type;

  @Column(name = "int_value")
  private Long int_value;

  @Column(name = "float_value")
  private Float float_value;

  @Column(name = "string_value", columnDefinition = "varchar")
  private String string_value;

  public HyperparameterValuesConfigBlob toProto() {
    return HyperparameterValuesConfigBlob.newBuilder()
        .setIntValue(this.int_value)
        .setFloatValue(this.float_value)
        .setStringValue(this.string_value)
        .build();
  }
}
