// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.swagger._public.modeldb.versioning.model

import scala.util.Try

import net.liftweb.json._

import ai.verta.swagger._public.modeldb.versioning.model.WorkspaceTypeEnumWorkspaceType._
import ai.verta.swagger.client.objects._

case class VersioningListCommitBlobsRequestResponse (
  blobs: Option[List[VersioningBlobExpanded]] = None,
  total_records: Option[String] = None
) extends BaseSwagger {
  def toJson(): JValue = VersioningListCommitBlobsRequestResponse.toJson(this)
}

object VersioningListCommitBlobsRequestResponse {
  def toJson(obj: VersioningListCommitBlobsRequestResponse): JObject = {
    new JObject(
      List[Option[JField]](
        obj.blobs.map(x => JField("blobs", ((x: List[VersioningBlobExpanded]) => JArray(x.map(((x: VersioningBlobExpanded) => VersioningBlobExpanded.toJson(x)))))(x))),
        obj.total_records.map(x => JField("total_records", JString(x)))
      ).flatMap(x => x match {
        case Some(y) => List(y)
        case None => Nil
      })
    )
  }

  def fromJson(value: JValue): VersioningListCommitBlobsRequestResponse =
    value match {
      case JObject(fields) => {
        val fieldsMap = fields.map(f => (f.name, f.value)).toMap
        VersioningListCommitBlobsRequestResponse(
          // TODO: handle required
          blobs = fieldsMap.get("blobs").map((x: JValue) => x match {case JArray(elements) => elements.map(VersioningBlobExpanded.fromJson); case _ => throw new IllegalArgumentException(s"unknown type ${x.getClass.toString}")}),
          total_records = fieldsMap.get("total_records").map(JsonConverter.fromJsonString)
        )
      }
      case _ => throw new IllegalArgumentException(s"unknown type ${value.getClass.toString}")
    }
}
