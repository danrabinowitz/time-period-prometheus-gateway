# Time Period Prometheus Gateway

Sometimes there is a need for a prometheus metric showing how much a value has increased for the current calendar month. For example, Comcast limits network usage to 1 terrabyte per calendar month. This can't be tracked natively in prometheus.

This gateway (exporter?) exposes metrics based on the "current month" and can easily be extended to "current day" or other time periods.
