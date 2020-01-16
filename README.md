# Time Period Prometheus Gateway

Sometimes there is a need for a prometheus metric showing how much a value has increased for the current calendar month. For example, Comcast limits network usage to 1 terrabyte per calendar month. This can't be tracked natively in prometheus.

This gateway (exporter?) exposes metrics based on the "current month" and can easily be extended to "current day" or other time periods.

There's been interest in this functionality:
https://www.robustperception.io/monthly-reporting-with-prometheus-and-python
https://community.grafana.com/t/displaying-totals-for-current-day-week-month-not-trailing/7721

This can be done in grafana via something like $__interval time selectors in grafana and a "this month" time selector. But that doesn't help with alerts.
