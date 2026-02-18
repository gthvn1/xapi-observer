**xapi-observer** is an open-source observability tool for the Xen XAPI toolstack.
It parses XAPI traces (Zipkin JSON), converts them to OpenTelemetry, and integrates with Grafana,
Tempo, and Loki for scalable trace and log analysis.


---

# Example of traces
```json
[
  {
    "tags": {
      "xs.xapi.task.id": "D:a53289ce0d41",
      "xs.xapi.session.track.id": "ed027918fda74f47eec4b3ac1de88018",
      "xs.pool.uuid": "0496aa55-4e6f-363d-6578-4ab872ddfb9e",
      "xs.observer.name": "local observer",
      "xs.host.uuid": "cb043a21-9c51-4418-a1e8-e7079ac6dc1a",
      "xs.host.name": "xcp-gtn-ip13",
      "service.name": "xapi"
    },
    "annotations": [],
    "localEndpoint": {
      "serviceName": "xapi"
    },
    "kind": "SERVER",
    "duration": 2,
    "timestamp": 1751469555858327,
    "name": "TaskHelper.get_name",
    "parentId": "873b73d304891032",
    "traceId": "0a07bfbf3925e3d4776995171a9222cd",
    "id": "d4bcd152c96bd105"
  },
  {
    "tags": {
      "xs.xapi.task.id": "D:31bcb54f6757",
      "xs.pool.uuid": "0496aa55-4e6f-363d-6578-4ab872ddfb9e",
      "xs.observer.name": "local observer",
      "xs.host.uuid": "cb043a21-9c51-4418-a1e8-e7079ac6dc1a",
      "xs.host.name": "xcp-gtn-ip13",
      "service.name": "xapi"
    },
    "annotations": [],
    "localEndpoint": {
      "serviceName": "xapi"
    },
    "duration": 2,
    "timestamp": 1751469555858210,
    "name": "TaskHelper.destroy",
    "traceId": "72aa2c3749bd41b16b9fd1b1d47b782a",
    "id": "fcdd53da8eb51fee"
  },
  {
    "tags": {
      "xs.xapi.task.uuid": "00000000-0000-0000-0000-000000000000",
      "xs.xapi.task.origin": "internal",
      "xs.xapi.task.name": "session_check",
      "xs.xapi.task.id": "D:31bcb54f6757",
      "xs.pool.uuid": "0496aa55-4e6f-363d-6578-4ab872ddfb9e",
      "xs.observer.name": "local observer",
      "xs.host.uuid": "cb043a21-9c51-4418-a1e8-e7079ac6dc1a",
      "xs.host.name": "xcp-gtn-ip13",
      "service.name": "xapi"
    },
    "annotations": [],
    "localEndpoint": {
      "serviceName": "xapi"
    },
    "duration": 82,
    "timestamp": 1751469555858125,
    "name": "session_check",
    "traceId": "8ba5b078e801c1dc85dc38e32bf39dcc",
    "id": "5fb761b593cbd395"
  },
  ...
```
