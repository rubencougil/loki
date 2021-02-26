# 📶 Loki/Grafana Docker + Sample Go App

Logging System built using **Loki** . It uses **Grafana** as the UI.

![Loki](https://user-images.githubusercontent.com/1073799/109333973-644ea000-7860-11eb-8c8a-65f4bb7ede1f.jpg)

## How to run it?

To spin up needed services

`docker-compose up -d`

Once done, the Sample Go App will automatically send some logs to Loki. Can be visualized in **Grafana** at:

```
http://localhost:3000
User: admin
Pass: foobar
```
