# RLF

RLF (Retrieve Log Files)

Welcome to RLF! 🧑‍🚀

Retrieve Log Files aka RLF is a service which helps us to fetch logs from with an easy and fast way.

### <u>Api</u>

The Api exposes the undermentioned endpoints:
Endpoint | HTTP Method | Parameters |Functionality |
--- | --- | -- | -- |
<b>/logs<b> | GET | report_type, datetime| Retrieve's a logs file based on provided parameters |

### <u>How to run:</u>

`go run .`

### Docker

To build the rlf image
`docker build -t rlf .`

To run the Docker container
`docker run -p 8080:8080 rlf`
