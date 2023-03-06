# Launchpad
_By Elon Society of Computing_

Launchpad is the the api center for the community at Elon University.

## Observability and Monitoring

This project uses the following technologies to enable observability and monitoring:
- [Grafana](https://grafana.com)
- [Loki](https://grafana.com/loki)
    - Logging in the backend provided by [Logrus](https://github.com/sirupsen/logrus)
- [Prometheus](https://prometheus.io/) 
    - (Though in the future we might consider [VictoriaMetrics](victoriametrics.com)!)
- [statsd](https://github.com/statsd/statsd)