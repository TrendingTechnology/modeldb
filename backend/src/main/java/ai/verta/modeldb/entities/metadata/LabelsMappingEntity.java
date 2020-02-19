package ai.verta.modeldb.entities.metadata;

import java.io.Serializable;
import java.util.Objects;
import javax.persistence.Column;
import javax.persistence.Embeddable;
import javax.persistence.EmbeddedId;
import javax.persistence.Entity;
import javax.persistence.Table;

@Entity
@Table(name = "labels_mapping")
public class LabelsMappingEntity {
  public LabelsMappingEntity() {}

  public LabelsMappingEntity(String entityHash, Integer entityType, String label) {
    this.id = new LabelMappingId(entityHash, entityType, label);
  }

  @EmbeddedId private LabelMappingId id;
}

@Embeddable
class LabelMappingId implements Serializable {

  @Column(name = "label", length = 50)
  private String label;

  @Column(name = "entity_type", length = 50)
  private Integer entity_type;

  @Column(name = "entity_hash", nullable = false, columnDefinition = "varchar", length = 32)
  private String entity_hash;

  public LabelMappingId(String entityHash, Integer entityType, String label) {
    this.entity_hash = entityHash;
    this.entity_type = entityType;
    this.label = label;
  }

  private LabelMappingId() {}

  public String getEntity_hash() {
    return entity_hash;
  }

  public Integer getEntity_type() {
    return entity_type;
  }

  public String getLabel() {
    return label;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (!(o instanceof LabelMappingId)) return false;
    LabelMappingId that = (LabelMappingId) o;
    return Objects.equals(getEntity_hash(), that.getEntity_hash())
        && Objects.equals(getEntity_type(), that.getEntity_type())
        && Objects.equals(getLabel(), that.getLabel());
  }

  @Override
  public int hashCode() {
    return Objects.hash(getEntity_hash(), getEntity_type(), getLabel());
  }
}
