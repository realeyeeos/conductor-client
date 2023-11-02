# conductor-client

1. 创建client
	client := sdk.NewClient("http://192.168.4.245:31500")
2. 创建工作流
	resp, err := client.CreateWorkFlow(data)
	data为工作流定义的json文件(需转为字符串格式)
	json格式示例：
{
  "name": "eos_workflow_test",
  "description": "EOS Workflow Test",
  "version": 1,
  "tasks": [
    {
      "name": "get_example_occupation",
      "taskReferenceName": "get_example_occupation",
      "description": null,
      "inputParameters": {
        "http_request": {
          "method": "GET",
          "uri": "http://192.168.2.168:8080/occupation"
        }
      },
      "type": "HTTP",
      "dynamicTaskNameParam": null,
      "caseValueParam": null,
      "caseExpression": null,
      "scriptExpression": null,
      "dynamicForkJoinTasksParam": null,
      "dynamicForkTasksParam": null,
      "dynamicForkTasksInputParamName": null,
      "startDelay": 0,
      "subWorkflowParam": null,
      "sink": null,
      "optional": false,
      "taskDefinition": null,
      "rateLimited": null,
      "asyncComplete": false,
      "loopCondition": null,
      "retryCount": null,
      "evaluatorType": null,
      "expression": null
    },
    {
      "name": "get_example_country",
      "taskReferenceName": "get_example_country",
      "description": null,
      "inputParameters": {
        "http_request": {
          "method": "GET",
          "uri": "http://192.168.2.168:8080/country"
        }
      },
      "type": "HTTP",
      "dynamicTaskNameParam": null,
      "caseValueParam": null,
      "caseExpression": null,
      "scriptExpression": null,
      "dynamicForkJoinTasksParam": null,
      "dynamicForkTasksParam": null,
      "dynamicForkTasksInputParamName": null,
      "startDelay": 0,
      "subWorkflowParam": null,
      "sink": null,
      "optional": false,
      "taskDefinition": null,
      "rateLimited": null,
      "asyncComplete": false,
      "loopCondition": null,
      "retryCount": null,
      "evaluatorType": null,
      "expression": null
    }
  ],
  "inputParameters": [],
  "outputParameters": {},
  "failureWorkflow": null,
  "schemaVersion": 2,
  "restartable": true,
  "workflowStatusListenerEnabled": false,
  "ownerEmail": null,
  "timeoutPolicy": "TIME_OUT_WF",
  "timeoutSeconds": 600,
  "variables": {},
  "inputTemplate": {}
}
3. 启动工作流
	resp, err := client.RunWorkFlow(data)
	data为工作流输入的json文件(需转为字符串格式)
	json格式示例：
{
  "name":"kitchensink1104", //工作流名称
  "version":"1", //工作流版本
  "input":{ //工作流输入
    "test":"sssmsmqkjammamam"
  }
}