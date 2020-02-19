// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.swagger._public.modeldb.model

import ai.verta.swagger._public.modeldb.model.ArtifactTypeEnumArtifactType._
import ai.verta.swagger._public.modeldb.model.OperatorEnumOperator._
import ai.verta.swagger._public.modeldb.model.TernaryEnumTernary._
import ai.verta.swagger._public.modeldb.model.ValueTypeEnumValueType._
import ai.verta.swagger._public.modeldb.model.ProtobufNullValue._

case class ModeldbCreateExperimentRun (
  id: Option[String],
  projectId: Option[String],
  experimentId: Option[String],
  name: Option[String],
  description: Option[String],
  dateCreated: Option[String],
  dateUpdated: Option[String],
  startTime: Option[String],
  endTime: Option[String],
  codeVersion: Option[String],
  codeVersionSnapshot: Option[ModeldbCodeVersion],
  parentId: Option[String],
  tags: Option[List[String]],
  attributes: Option[List[CommonKeyValue]],
  hyperparameters: Option[List[CommonKeyValue]],
  artifacts: Option[List[ModeldbArtifact]],
  datasets: Option[List[ModeldbArtifact]],
  metrics: Option[List[CommonKeyValue]],
  observations: Option[List[ModeldbObservation]],
  features: Option[List[ModeldbFeature]]
)
