// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.swagger._public.modeldb.versioning.model

import scala.util.Try

import net.liftweb.json._

import ai.verta.swagger._public.modeldb.versioning.model.WorkspaceTypeEnumWorkspaceType._
import ai.verta.swagger.client.objects._

case class VersioningDatasetDiff (
  s3: Option[VersioningS3DatasetDiff] = None,
  path: Option[VersioningPathDatasetDiff] = None
) extends BaseSwagger {
  def toJson(): JValue = VersioningDatasetDiff.toJson(this)
}

object VersioningDatasetDiff {
  def toJson(obj: VersioningDatasetDiff): JObject = {
    new JObject(
      List[Option[JField]](
        obj.s3.map(x => JField("s3", ((x: VersioningS3DatasetDiff) => VersioningS3DatasetDiff.toJson(x))(x))),
        obj.path.map(x => JField("path", ((x: VersioningPathDatasetDiff) => VersioningPathDatasetDiff.toJson(x))(x)))
      ).flatMap(x => x match {
        case Some(y) => List(y)
        case None => Nil
      })
    )
  }

  def fromJson(value: JValue): VersioningDatasetDiff =
    value match {
      case JObject(fields) => {
        val fieldsMap = fields.map(f => (f.name, f.value)).toMap
        VersioningDatasetDiff(
          // TODO: handle required
          s3 = fieldsMap.get("s3").map(VersioningS3DatasetDiff.fromJson),
          path = fieldsMap.get("path").map(VersioningPathDatasetDiff.fromJson)
        )
      }
      case _ => throw new IllegalArgumentException(s"unknown type ${value.getClass.toString}")
    }
}
