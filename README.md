# Forex-dashboard

ForexDashboard is an application that can generate price data for currency pairs (e.g. EURUSD exchange rate) and show the generated data real-time on a candlestick chart.
<br/>
I made this small project to learn Go and microservice architecture. Every service has it's own docker container.
RabbitMQ is used for communication between services. 

TODOS:
- [x] live reload in docker on code change
- [x] graceful shutdown on server
- [ ] fix model properties (use Meta instead of CurrencyPair)
- [ ] file service with download/upload feature
- [ ] healtcheck api
- [ ] create proper controllers in every service
- [ ] create proper mock response for the chart
- [ ] standardize events (create service for messaging?)
- [ ] create common send/recieve functions so all service can use them

## Good to have
- [ ] event journal in db
- [ ] mock data points in db
- [ ] jwt authentication
- [ ] service template for creating new services
