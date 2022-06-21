> ### ⚠️ Deprecated
> This tool has been deprecated in favor of [cos-tool](https://github.com/canonical/cos-tool/),
> which in addition to being able to everything this tool does, also is able to validate alert rules and expressions,
> as well as operate on logql expressions.

# promql-transform

Transforms PromQL expressions on the fly

## Usage

Given the expression
```
job:request_latency_seconds:mean5m{job=\"myjob\"} > 0.5
```

Running

```bash
$ ./promql-transform \
    --label-matcher juju_model=lma \
    --label-matcher juju_model_uuid=12345 \
    --label-matcher juju_application=proxy \
    --label-matcher juju_unit=proxy/1 \
    "job:request_latency_seconds:mean5m{job=\"myjob\"} > 0.5"
```

Would output

```
job:request_latency_seconds:mean5m{job="myjob",juju_application="proxy",juju_model="lma",juju_model_uuid="12345",juju_unit="proxy/1"} > 0.5
```
