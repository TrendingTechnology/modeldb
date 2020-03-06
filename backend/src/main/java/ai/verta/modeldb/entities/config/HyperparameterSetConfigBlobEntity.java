package ai.verta.modeldb.entities.config;

import ai.verta.modeldb.ModelDBException;
import ai.verta.modeldb.versioning.ContinuousHyperparameterSetConfigBlob;
import ai.verta.modeldb.versioning.DiscreteHyperparameterSetConfigBlob;
import ai.verta.modeldb.versioning.HyperparameterSetConfigBlob;
import ai.verta.modeldb.versioning.HyperparameterSetConfigBlob.ValueCase;
import ai.verta.modeldb.versioning.HyperparameterValuesConfigBlob;
import io.grpc.Status.Code;
import java.util.ArrayList;
import java.util.List;
import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.OneToMany;
import javax.persistence.OneToOne;
import javax.persistence.Table;

@Entity
@Table(name = "hyperparameter_set_config_blob")
public class HyperparameterSetConfigBlobEntity {
  private HyperparameterSetConfigBlobEntity() {}

  public HyperparameterSetConfigBlobEntity(
      String blobHash, String commitHash, HyperparameterSetConfigBlob hyperparameterSetConfigBlob)
      throws ModelDBException {
    this.blob_hash = blobHash;
    this.name = hyperparameterSetConfigBlob.getName();
    this.value_type = hyperparameterSetConfigBlob.getValueCase().getNumber();

    switch (hyperparameterSetConfigBlob.getValueCase()) {
      case CONTINUOUS:
        ContinuousHyperparameterSetConfigBlob configBlobContinuous =
            hyperparameterSetConfigBlob.getContinuous();
        this.interval_begin_hash =
            new HyperparameterElementConfigBlobEntity(
                blobHash, commitHash, null, configBlobContinuous.getIntervalBegin());
        this.interval_end_hash =
            new HyperparameterElementConfigBlobEntity(
                blobHash, commitHash, null, configBlobContinuous.getIntervalEnd());
        this.interval_step_hash =
            new HyperparameterElementConfigBlobEntity(
                blobHash, commitHash, null, configBlobContinuous.getIntervalStep());
        break;
      case DISCRETE:
        DiscreteHyperparameterSetConfigBlob configBlobContinuousDiscrete =
            hyperparameterSetConfigBlob.getDiscrete();
        List<HyperparameterValuesConfigBlob> valuesConfigBlobs =
            configBlobContinuousDiscrete.getValuesList();

        for (HyperparameterValuesConfigBlob hyperparameterValuesConfigBlob : valuesConfigBlobs) {
          this.hyperparameterSetConfigElementMapping.add(
              new HyperparameterElementConfigBlobEntity(
                  blobHash, commitHash, null, hyperparameterValuesConfigBlob));
        }
        break;
      case VALUE_NOT_SET:
      default:
        throw new ModelDBException(
            "Invalid value found for HyperparameterSetConfigBlob", Code.INVALID_ARGUMENT);
    }
  }

  @Id
  @Column(name = "blob_hash", columnDefinition = "varchar", length = 64, nullable = false)
  private String blob_hash;

  @Column(name = "name", columnDefinition = "varchar")
  private String name;

  @Column(name = "value_type")
  private Integer value_type;

  @OneToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "interval_begin_hash")
  private HyperparameterElementConfigBlobEntity interval_begin_hash;

  @OneToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "interval_end_hash")
  private HyperparameterElementConfigBlobEntity interval_end_hash;

  @OneToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "interval_step_hash")
  private HyperparameterElementConfigBlobEntity interval_step_hash;

  @OneToMany
  @JoinTable(
      name = "hyperparameter_discrete_set_element_mapping",
      joinColumns = {@JoinColumn(name = "set_hash")},
      inverseJoinColumns = {@JoinColumn(name = "element_hash")})
  private List<HyperparameterElementConfigBlobEntity> hyperparameterSetConfigElementMapping =
      new ArrayList<>();

  public HyperparameterSetConfigBlob toProto() throws ModelDBException {
    HyperparameterSetConfigBlob.Builder builder =
        HyperparameterSetConfigBlob.newBuilder().setName(this.name);
    ValueCase valueCase = ValueCase.forNumber(this.value_type);
    if (valueCase == null) {
      throw new ModelDBException(
          "Invalid value found for HyperparameterSetConfigBlob", Code.INVALID_ARGUMENT);
    }

    switch (valueCase) {
      case CONTINUOUS:
        builder.setContinuous(
            ContinuousHyperparameterSetConfigBlob.newBuilder()
                .setIntervalBegin(this.interval_begin_hash.toProto())
                .setIntervalEnd(this.interval_end_hash.toProto())
                .setIntervalStep(this.interval_step_hash.toProto())
                .build());
        break;
      case DISCRETE:
        break;
      case VALUE_NOT_SET:
      default:
        throw new ModelDBException(
            "Invalid value found for HyperparameterSetConfigBlob", Code.INVALID_ARGUMENT);
    }
    return builder.build();
  }
}
