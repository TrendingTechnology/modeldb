# -*- coding: utf-8 -*-

from __future__ import print_function

from .._protos.public.modeldb.versioning import VersioningService_pb2 as _VersioningService

from .. import _utils
from . import commit


class Repository(object):
    def __init__(self, conn, id_):
        self._conn = conn

        self.id = id_

    def __repr__(self):
        return "<Repository \"{}\">".format(self.name)

    @property
    def _endpoint_prefix(self):
        return "{}://{}/api/v1/modeldb/versioning/repositories/{}".format(
            self._conn.scheme,
            self._conn.socket,
            self.id,
        )

    @property
    def name(self):
        response = _utils.make_request("GET", self._endpoint_prefix, self._conn)
        _utils.raise_for_http_error(response)

        response_msg = _utils.json_to_proto(response.json(),
                                            _VersioningService.GetRepositoryRequest.Response)
        return response_msg.repository.name

    @property
    def workspace(self):
        raise NotImplementedError

    @classmethod
    def _create(cls, conn, name, workspace):
        msg = _VersioningService.Repository()
        msg.name = name

        data = _utils.proto_to_json(msg)
        endpoint = "{}://{}/api/v1/modeldb/versioning/workspaces/{}/repositories".format(
            conn.scheme,
            conn.socket,
            workspace,
        )
        response = _utils.make_request("POST", endpoint, conn, json=data)
        _utils.raise_for_http_error(response)

        response_msg = _utils.json_to_proto(response.json(),
                                            _VersioningService.SetRepository.Response)
        return cls(conn, response_msg.repository.id)

    @classmethod
    def _get(cls, conn, name=None, workspace=None, id_=None):
        if name and workspace and not id_:
            endpoint = "{}://{}/api/v1/modeldb/versioning/workspaces/{}/repositories/{}".format(
                conn.scheme,
                conn.socket,
                workspace,
                name,
            )
        elif not name and not workspace and id_:
            endpoint = "{}://{}/api/v1/modeldb/versioning/repositories/{}".format(
                conn.scheme,
                conn.socket,
                id_,
            )
        else:
            raise RuntimeError("the Client has encountered an error;"
                               " please notify the Verta development team")
        response = _utils.make_request("GET", endpoint, conn)

        if not response.ok:
            if ((response.status_code == 403 and response.json()['code'] == 7)
                    or (response.status_code == 404 and response.json()['code'] == 5)):
                return None
            else:
                _utils.raise_for_http_error(response)

        response_msg = _utils.json_to_proto(response.json(),
                                            _VersioningService.GetRepositoryRequest.Response)
        return cls(conn, response_msg.repository.id)

    def new_commit(self, parents=None):
        if parents is None:
            parents = []

        parent_ids = []
        for i, parent in enumerate(parents):
            if not isinstance(parent, commit.Commit):
                raise TypeError("`parents` must only contain Commits, not {}".format(type(parent)))
            if parent.id is None:
                raise ValueError("parent at index {} does not have an ID;"
                                 " please save it first".format(i))

            parent_ids.append(parent.id)

        return commit.Commit(self._conn, self.id, parent_ids)

    def get_commit(self, tag=None, id=None):
        if tag is not None and id is not None:
            raise ValueError("cannot specify both `tag` and `id`")
        elif tag is not None:
            msg = _VersioningService.GetTagRequest()
            endpoint = "{}://{}/api/v1/modeldb/versioning/repositories/{}/tags/{}".format(
                self._conn.scheme,
                self._conn.socket,
                self.id,
                tag,
            )
        elif id is not None:
            msg = _VersioningService.GetCommitRequest()
            endpoint = "{}://{}/api/v1/modeldb/versioning/repositories/{}/commits/{}".format(
                self._conn.scheme,
                self._conn.socket,
                self.id,
                id,
            )
        else:
            raise ValueError("must specify either `tag` or `id`")
        response = _utils.make_request("GET", endpoint, self._conn)
        _utils.raise_for_http_error(response)

        response_msg = _utils.json_to_proto(response.json(), msg.Response)
        return commit.Commit._from_id(self._conn, self.id, response_msg.commit.commit_sha)
