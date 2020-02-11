import six

import itertools
import os
import shutil

import requests

import pytest
import utils

import verta
import json


KWARGS = {
    'desc': [None, "A test."],
    'tags': [None, ['test']],
    'attrs': [None, {'is_test': True}],
}
KWARGS_COMBOS = [dict(zip(KWARGS.keys(), values))
                 for values
                 in itertools.product(*KWARGS.values())
                 if values.count(None) != len(values)]


class TestClient:
    @pytest.mark.oss
    def test_no_auth(self, host):
        client = verta.Client(host)

        # it's just been revoked
        client._conn.auth = None

        assert client.set_project()
        utils.delete_project(client.proj.id, client._conn)

    @pytest.mark.skipif('VERTA_EMAIL' not in os.environ or 'VERTA_DEV_KEY' not in os.environ, reason="insufficient Verta credentials")
    def test_verta_https(self):
        hosts = [
            "dev.verta.ai",
            "app.verta.ai",
            "sandbox.app.verta.ai",
        ]

        for host in hosts:
            # https by default
            conn = verta.Client(host)._conn
            assert conn.scheme == "https"
            assert conn.scheme == conn.auth['Grpc-Metadata-scheme']

            # http if provided
            conn = verta.Client("http://{}".format(host))._conn
            assert conn.scheme == "http"
            assert conn.scheme == conn.auth['Grpc-Metadata-scheme']

            # https if provided
            conn = verta.Client("https://{}".format(host))._conn
            assert conn.scheme == "https"
            assert conn.scheme == conn.auth['Grpc-Metadata-scheme']

    @pytest.mark.skipif('VERTA_EMAIL' not in os.environ or 'VERTA_DEV_KEY' not in os.environ, reason="insufficient Verta credentials")
    def test_config_file(self):
        SOCKET = "sandbox.app.verta.ai"
        EMAIL_KEY = "VERTA_EMAIL"
        DEV_KEY_KEY = "VERTA_DEV_KEY"
        email = os.environ[EMAIL_KEY]
        dev_key = os.environ[DEV_KEY_KEY]
        del os.environ[EMAIL_KEY]
        del os.environ[DEV_KEY_KEY]
        PROJECT_NAME = "test_project"
        DATASET_NAME = "test_dataset"
        EXPERIMENT_NAME = "test_experiment"
        CONFIG_FILENAME = "verta_config.json"
        open(CONFIG_FILENAME, 'w').write(json.dumps({"email": email, "dev_key": dev_key,
                                                         "host": SOCKET,
                                                         "project": PROJECT_NAME,
                                                         "dataset": DATASET_NAME,
                                                         "experiment": EXPERIMENT_NAME}))
        client = verta.Client()
        conn = client._conn
        assert conn.socket == SOCKET
        assert conn.auth['Grpc-Metadata-email'] == email
        assert conn.auth['Grpc-Metadata-developer_key'] == dev_key
        proj = client.set_project()
        assert proj.name == PROJECT_NAME
        expt = client.set_experiment()
        assert expt.name == EXPERIMENT_NAME
        dataset = client.set_dataset()
        assert dataset.name == DATASET_NAME
        utils.delete_project(proj.id, conn)
        utils.delete_datasets([dataset.id], conn)
        os.environ[EMAIL_KEY] = email
        os.environ[DEV_KEY_KEY] = dev_key
        os.remove(CONFIG_FILENAME)

    def test_else_http(self):
        # test hosts must not redirect http to https
        hosts = [
            "www.google.com",
        ]

        for host in hosts:
            # http by default
            try:
                verta.Client(host, max_retries=0)
            except requests.HTTPError as e:
                assert e.request.url.split(':', 1)[0] == "http"
            else:
                raise RuntimeError("faulty test; expected error")

            # http if provided
            try:
                verta.Client("http://{}".format(host), max_retries=0)
            except requests.HTTPError as e:
                assert e.request.url.split(':', 1)[0] == "http"
            else:
                raise RuntimeError("faulty test; expected error")

            # https if provided
            try:
                verta.Client("https://{}".format(host), max_retries=0)
            except requests.HTTPError as e:
                assert e.request.url.split(':', 1)[0] == "https"
            else:
                raise RuntimeError("faulty test; expected error")


class TestEntities:
    def test_cache(self, client, strs):
        entities = (
            client.set_project(),
            client.set_experiment(),
            client.set_experiment_run(),
        )

        for entity in entities:
            filename = strs[0]
            filepath = os.path.join(verta.client._CACHE_DIR, filename)
            contents = six.ensure_binary(strs[1])

            assert not os.path.isfile(filepath)
            assert not entity._get_cached(filename)

            try:
                assert entity._cache(filename, contents) == filepath

                assert os.path.isfile(filepath)
                assert entity._get_cached(filename)

                with open(filepath, 'rb') as f:
                    assert f.read() == contents
            finally:
                shutil.rmtree(verta.client._CACHE_DIR, ignore_errors=True)


class TestProject:
    def test_set_project_warning(self, client):
        """setting Project by name with desc, tags, and/or attrs raises warning"""
        proj = client.set_project()

        for kwargs in KWARGS_COMBOS:
            with pytest.warns(UserWarning):
                client.set_project(proj.name, **kwargs)

    def test_create(self, client):
        assert client.set_project()

        assert client.proj is not None

    def test_get_by_name(self, client):
        proj = client.set_project()

        client.set_project()  # in case get erroneously fetches latest

        assert proj.id == client.set_project(proj.name).id

    def test_get_by_id(self, client):
        proj = client.set_project()

        client.set_project()  # in case get erroneously fetches latest
        client.proj = None

        assert proj.id == client.set_project(id=proj.id).id

    def test_get_nonexistent_id(self, client):
        with pytest.raises(ValueError):
            client.set_project(id="nonexistent_id")


class TestExperiment:
    def test_set_experiment_warning(self, client):
        """setting Experiment by name with desc, tags, and/or attrs raises warning"""
        client.set_project()
        expt = client.set_experiment()

        for kwargs in KWARGS_COMBOS:
            with pytest.warns(UserWarning):
                client.set_experiment(expt.name, **kwargs)

    def test_create(self, client):
        client.set_project()
        assert client.set_experiment()

        assert client.expt is not None

    def test_get_by_name(self, client):
        client.set_project()
        expt = client.set_experiment()

        client.set_experiment()  # in case get erroneously fetches latest

        assert expt.id == client.set_experiment(expt.name).id

    def test_get_by_id(self, client):
        proj = client.set_project()
        expt = client.set_experiment()

        client.set_experiment()  # in case get erroneously fetches latest
        client.proj = client.expt = None

        assert expt.id == client.set_experiment(id=expt.id).id
        assert proj.id == client.proj.id

    def test_get_nonexistent_id(self, client):
        with pytest.raises(ValueError):
            client.set_experiment(id="nonexistent_id")


class TestExperimentRun:
    def test_set_experiment_run_warning(self, client):
        """setting ExperimentRun by name with desc, tags, and/or attrs raises warning"""
        client.set_project()
        client.set_experiment()
        expt_run = client.set_experiment_run()

        for kwargs in KWARGS_COMBOS:
            with pytest.warns(UserWarning):
                client.set_experiment_run(expt_run.name, **kwargs)

    def test_create(self, client):
        client.set_project()
        client.set_experiment()

        assert client.set_experiment_run()

    def test_get_by_name(self, client):
        client.set_project()
        client.set_experiment()
        run = client.set_experiment_run()
        client.set_experiment_run()  # in case get erroneously fetches latest

        assert run.id == client.set_experiment_run(run.name).id

    def test_get_by_id(self, client):
        proj = client.set_project()
        expt = client.set_experiment()
        expt_run = client.set_experiment_run()

        client.set_experiment_run()  # in case get erroneously fetches latest
        client.proj = client.expt = None

        assert expt_run.id == client.set_experiment_run(id=expt_run.id).id
        assert proj.id == client.proj.id
        assert expt.id == client.expt.id

    def test_get_nonexistent_id(self, client):
        with pytest.raises(ValueError):
            client.set_experiment_run(id="nonexistent_id")

    def test_clone(self, experiment_run):
        expt_run = experiment_run
        expt_run._conf.use_git = False
        expt_run.log_hyperparameters({"hpp1" : 1, "hpp2" : 2, "hpp3" : "hpp3"})
        expt_run.log_metrics({"metric1" : 0.5, "metric2" : 0.6})
        expt_run.log_tags(["tag1", "tag2"])
        expt_run.log_attributes({"attr1" : 10, "attr2" : {"abc": 1}})
        expt_run.log_artifact("my-artifact", "README.md")
        expt_run.log_code()

        # set various things in the run
        new_run_no_art = expt_run.clone()
        new_run_art_only = expt_run.clone(copy_artifacts=True)
        new_run_art_code = expt_run.clone(copy_artifacts=True, copy_code_version=True)
        new_run_art_code_data = expt_run.clone(copy_artifacts=True,
            copy_code_version=True, copy_datasets=True)

        old_run_msg = expt_run._get(expt_run._conn, _expt_run_id=expt_run.id)
        new_run_no_art_msg = new_run_no_art._get(new_run_no_art._conn, _expt_run_id=new_run_no_art.id)
        new_run_art_only_msg = new_run_art_only._get(new_run_art_only._conn, _expt_run_id=new_run_art_only.id)
        new_run_art_code_msg = new_run_art_code._get(new_run_art_code._conn, _expt_run_id=new_run_art_code.id)
        new_run_art_code_data_msg = new_run_art_code_data._get(new_run_art_code_data._conn,
            _expt_run_id=new_run_art_code_data.id)

        # ensure basic data is the same
        assert expt_run.id != new_run_no_art_msg.id
        assert old_run_msg.description == new_run_no_art_msg.description
        assert old_run_msg.tags == new_run_no_art_msg.tags
        assert old_run_msg.metrics == new_run_no_art_msg.metrics
        assert old_run_msg.hyperparameters == new_run_no_art_msg.hyperparameters
        assert old_run_msg.observations == new_run_no_art_msg.observations

        assert old_run_msg.artifacts == new_run_art_only_msg.artifacts
        assert old_run_msg.code_version_snapshot != new_run_art_only_msg.code_version_snapshot
        assert old_run_msg.artifacts != new_run_no_art_msg.artifacts

        assert old_run_msg.code_version_snapshot == new_run_art_code_msg.code_version_snapshot

        assert old_run_msg.datasets == new_run_art_code_data_msg.datasets

class TestExperimentRuns:
    def test_getitem(self, client):
        client.set_project()
        expt = client.set_experiment()

        # test for...in iteration
        local_run_ids = set(client.set_experiment_run().id for _ in range(3))
        assert local_run_ids == set(run.id for run in expt.expt_runs)

        # do it again
        local_run_ids.update(client.set_experiment_run().id for _ in range(3))
        assert local_run_ids == set(run.id for run in expt.expt_runs)

        # direct __getitem__
        assert expt.expt_runs[0].id in local_run_ids

    def test_len(self, client):
        client.set_project()
        expt = client.set_experiment()
        run_ids = [client.set_experiment_run().id for _ in range(3)]

        assert len(run_ids) == len(expt.expt_runs)

    def test_add(self, client):
        client.set_project()
        expt1 = client.set_experiment()
        local_expt1_run_ids = set(client.set_experiment_run().id for _ in range(3))
        expt2 = client.set_experiment()
        local_expt2_run_ids = set(client.set_experiment_run().id for _ in range(3))

        # simple concatenation
        assert local_expt1_run_ids | local_expt2_run_ids == set(run.id for run in expt1.expt_runs + expt2.expt_runs)

        # ignore duplicates
        assert local_expt1_run_ids == set(run.id for run in expt1.expt_runs + expt1.expt_runs)
