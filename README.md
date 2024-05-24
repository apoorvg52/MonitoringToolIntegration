# Monitoringtool/api with new relics in Golang
MELTs Integration for better management &amp; insights on the errors

## Description
This project aims to integrate MELTs (Metrics, Events, Logs, Trace) with New Relic's APIs to provide better management and insights into errors and system performance. It can be easily use by any company to send the data to new relics server.

## Features
- Integration with New Relic's Metrics API to track system metrics such as CPU usage, memory usage, and network traffic.
- Integration with New Relic's Events API to log custom events, such as user actions or system events.
- Integration with New Relic's Logs API to send log data for analysis and monitoring.
- Integration with New Relic's Trace API to trace requests through the system and identify performance bottlenecks.

## Requirements
- Go 1.20 or later

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/apoorvg52/MonitoringToolIntegration.git
   cd MonitoringToolIntegration

2. Configure your New Relic API key and other settings in the utils/constants file.
3. Run the project
    ```bash
    go run main.go  
4. Local debugging configuration for VSCode has been preconfigured. Simply navigate to the "Run & Debug" section.

5. Initiate the project in debug mode by selecting "Launch Project".

   
## Reference Screenshots
**API Supported**

Description: This screenshot shows the list of APIs supported by the integration with New Relic.


1.**`/uploadEvents`**
   Description: This endpoint handles the upload of custom event data to New Relic. This endpoint accepts a JSON file as input and compresses it to gzip format before sending it to New Relic's server for further processing and analysis.
   http://localhost:8080/uploadEvents
    a. **Request from Postman:**
    ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/b46b1057-490d-4951-9b8c-a03859d80920)
    ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/d23f4e44-7190-4173-850e-c52d23c9594b)

    b. **Response Received with Custom Data Points:**
    
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/02f23cb5-f703-44b0-a032-0ee48acde057)

    c. **Stats on New Relic's UI:**
![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/375e9661-ba73-413e-b4fa-f2fe2c6abba6)
    
    



2. **`/uploadMetric`**
   Description: This API endpoint is used to upload metrics data to New Relic.
   http://localhost:8080/uploadMetric

    a. **Request from Postman:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/06538b82-d00d-47c3-8a2e-5c79d657bb5e)

    b. **Response Received with Custom Data Points:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/a802c0f0-60ef-450e-8532-51f130cb7cd0)

    c. **Stats on New Relic's UI:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/53078968-18bd-4667-8bb9-1c4c18fc83e0)
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/c5585f04-868f-4b55-afab-98505bee5aec)


   
3. **`/uploadLogs`**
   Description: This API endpoint is employed to upload log data to New Relic for analysis and monitoring.
   http://localhost:8080/uploadLogs

    a. **Request from Postman:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/bd05401c-29b4-416f-a8aa-0f62853d2caf)

    b. **Response Received with Custom Data Points:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/cad8bfe3-b576-407c-9e3f-d28f3ece0fca)

    c. **Logs Data on New Relic's UI:**
   ![5 logs   metrics](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/ccad0f4c-215f-44b2-be02-38734d30ba78)


4. **`/uploadTrace`**
   Description: This API endpoint facilitates the uploading of trace data to New Relic for performance monitoring and analysis.
   http://localhost:8080/uploadTrace

    a. **Request from Postman:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/5b4f0bc4-f93e-4661-8efa-791b431313f3)

    b. **Response Received with Custom Data Points:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/e0c307e5-3c63-442c-a26a-4d5edbbfcd32)

    c. **Trace Data on New Relic's UI:**
   ![image](https://github.com/apoorvg52/MonitoringToolIntegration/assets/61585615/d3b99519-ed35-43c4-92cc-4be19aa31e64)

   
## Usage
- Use the provided functions to send metrics, events, logs, and traces to New Relic.
- Customize the integration to suit your specific monitoring and analysis needs.
- Refer to the New Relic documentation for details on using the Metrics, Events, Logs, and Trace APIs.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue if you encounter any problems or have suggestions for improvements.

## Acknowledgements
Thanks to the New Relic team for providing comprehensive APIs for monitoring and analysis.
