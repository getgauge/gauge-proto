# Protocol Documentation
<a name="top"/>

## Table of Contents
* [api.proto](#api.proto)
 * [APIMessage](#gauge.messages.APIMessage)
 * [ConceptInfo](#gauge.messages.ConceptInfo)
 * [ErrorResponse](#gauge.messages.ErrorResponse)
 * [ExtractConceptInfoRequest](#gauge.messages.ExtractConceptInfoRequest)
 * [ExtractConceptInfoResponse](#gauge.messages.ExtractConceptInfoResponse)
 * [FormatConceptHeadingRequest](#gauge.messages.FormatConceptHeadingRequest)
 * [FormatConceptHeadingResponse](#gauge.messages.FormatConceptHeadingResponse)
 * [GetAllConceptsRequest](#gauge.messages.GetAllConceptsRequest)
 * [GetAllConceptsResponse](#gauge.messages.GetAllConceptsResponse)
 * [GetAllSpecsRequest](#gauge.messages.GetAllSpecsRequest)
 * [GetAllSpecsResponse](#gauge.messages.GetAllSpecsResponse)
 * [GetAllStepsRequest](#gauge.messages.GetAllStepsRequest)
 * [GetAllStepsResponse](#gauge.messages.GetAllStepsResponse)
 * [GetInstallationRootRequest](#gauge.messages.GetInstallationRootRequest)
 * [GetInstallationRootResponse](#gauge.messages.GetInstallationRootResponse)
 * [GetLanguagePluginLibPathRequest](#gauge.messages.GetLanguagePluginLibPathRequest)
 * [GetLanguagePluginLibPathResponse](#gauge.messages.GetLanguagePluginLibPathResponse)
 * [GetProjectRootRequest](#gauge.messages.GetProjectRootRequest)
 * [GetProjectRootResponse](#gauge.messages.GetProjectRootResponse)
 * [GetStepValueRequest](#gauge.messages.GetStepValueRequest)
 * [GetStepValueResponse](#gauge.messages.GetStepValueResponse)
 * [PerformRefactoringRequest](#gauge.messages.PerformRefactoringRequest)
 * [PerformRefactoringResponse](#gauge.messages.PerformRefactoringResponse)
 * [APIMessage.APIMessageType](#gauge.messages.APIMessage.APIMessageType)
* [messages.proto](#messages.proto)
 * [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest)
 * [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest)
 * [ExecutionInfo](#gauge.messages.ExecutionInfo)
 * [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest)
 * [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse)
 * [KillProcessRequest](#gauge.messages.KillProcessRequest)
 * [Message](#gauge.messages.Message)
 * [ParameterPosition](#gauge.messages.ParameterPosition)
 * [RefactorRequest](#gauge.messages.RefactorRequest)
 * [RefactorResponse](#gauge.messages.RefactorResponse)
 * [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest)
 * [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest)
 * [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest)
 * [ScenarioInfo](#gauge.messages.ScenarioInfo)
 * [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest)
 * [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest)
 * [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest)
 * [SpecInfo](#gauge.messages.SpecInfo)
 * [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest)
 * [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest)
 * [StepInfo](#gauge.messages.StepInfo)
 * [StepNameRequest](#gauge.messages.StepNameRequest)
 * [StepNameResponse](#gauge.messages.StepNameResponse)
 * [StepNamesRequest](#gauge.messages.StepNamesRequest)
 * [StepNamesResponse](#gauge.messages.StepNamesResponse)
 * [StepValidateRequest](#gauge.messages.StepValidateRequest)
 * [StepValidateResponse](#gauge.messages.StepValidateResponse)
 * [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest)
 * [SuiteExecutionResult](#gauge.messages.SuiteExecutionResult)
 * [Message.MessageType](#gauge.messages.Message.MessageType)
* [spec.proto](#spec.proto)
 * [Fragment](#gauge.messages.Fragment)
 * [Parameter](#gauge.messages.Parameter)
 * [ProtoComment](#gauge.messages.ProtoComment)
 * [ProtoConcept](#gauge.messages.ProtoConcept)
 * [ProtoExecutionResult](#gauge.messages.ProtoExecutionResult)
 * [ProtoHookFailure](#gauge.messages.ProtoHookFailure)
 * [ProtoItem](#gauge.messages.ProtoItem)
 * [ProtoScenario](#gauge.messages.ProtoScenario)
 * [ProtoSpec](#gauge.messages.ProtoSpec)
 * [ProtoSpecResult](#gauge.messages.ProtoSpecResult)
 * [ProtoStep](#gauge.messages.ProtoStep)
 * [ProtoStepExecutionResult](#gauge.messages.ProtoStepExecutionResult)
 * [ProtoStepValue](#gauge.messages.ProtoStepValue)
 * [ProtoSuiteResult](#gauge.messages.ProtoSuiteResult)
 * [ProtoTable](#gauge.messages.ProtoTable)
 * [ProtoTableDrivenScenario](#gauge.messages.ProtoTableDrivenScenario)
 * [ProtoTableRow](#gauge.messages.ProtoTableRow)
 * [ProtoTags](#gauge.messages.ProtoTags)
 * [Fragment.FragmentType](#gauge.messages.Fragment.FragmentType)
 * [Parameter.ParameterType](#gauge.messages.Parameter.ParameterType)
 * [ProtoItem.ItemType](#gauge.messages.ProtoItem.ItemType)
* [Scalar Value Types](#scalar-value-types)

<a name="api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## api.proto

<a name="gauge.messages.APIMessage"/>
### APIMessage
A generic message composing of all possible operations.
/ One of the Request/Response fields will have value, depending on the MessageType set.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| messageType | APIMessage.APIMessageType | required | Type of API call being made |
| messageId | int64 | required | A unique id to represent this message. A response to the message should copy over this value./ This is used to synchronize messages &amp; responses |
| projectRootRequest | GetProjectRootRequest | optional | [GetProjectRootRequest](#gauge.messages.GetProjectRootRequest) |
| projectRootResponse | GetProjectRootResponse | optional | [GetProjectRootResponse](#gauge.messages.GetProjectRootResponse) |
| installationRootRequest | GetInstallationRootRequest | optional | [GetInstallationRootRequest](#gauge.messages.GetInstallationRootRequest) |
| installationRootResponse | GetInstallationRootResponse | optional | [GetInstallationRootResponse](#gauge.messages.GetInstallationRootResponse) |
| allStepsRequest | GetAllStepsRequest | optional | [GetAllStepsRequest](#gauge.messages.GetAllStepsRequest) |
| allStepsResponse | GetAllStepsResponse | optional | [GetAllStepsResponse](#gauge.messages.GetAllStepsResponse) |
| allSpecsRequest | GetAllSpecsRequest | optional | [GetAllSpecsRequest](#gauge.messages.GetAllSpecsRequest) |
| allSpecsResponse | GetAllSpecsResponse | optional | [GetAllSpecsResponse](#gauge.messages.GetAllSpecsResponse) |
| stepValueRequest | GetStepValueRequest | optional | [GetStepValueRequest](#gauge.messages.GetStepValueRequest) |
| stepValueResponse | GetStepValueResponse | optional | [GetStepValueResponse](#gauge.messages.GetStepValueResponse) |
| libPathRequest | GetLanguagePluginLibPathRequest | optional | [GetLanguagePluginLibPathRequest](#gauge.messages.GetLanguagePluginLibPathRequest) |
| libPathResponse | GetLanguagePluginLibPathResponse | optional | [GetLanguagePluginLibPathResponse](#gauge.messages.GetLanguagePluginLibPathResponse) |
| error | ErrorResponse | optional | [ErrorResponse](#gauge.messages.ErrorResponse) |
| allConceptsRequest | GetAllConceptsRequest | optional | [GetAllConceptsRequest](#gauge.messages.GetAllConceptsRequest) |
| allConceptsResponse | GetAllConceptsResponse | optional | [GetAllConceptsResponse](#gauge.messages.GetAllConceptsResponse) |
| performRefactoringRequest | PerformRefactoringRequest | optional | [PerformRefactoringRequest](#gauge.messages.PerformRefactoringRequest) |
| performRefactoringResponse | PerformRefactoringResponse | optional | [PerformRefactoringResponse](#gauge.messages.PerformRefactoringResponse) |
| extractConceptInfoRequest | ExtractConceptInfoRequest | optional | [ExtractConceptInfoRequest](#gauge.messages.ExtractConceptInfoRequest) |
| extractConceptInfoResponse | ExtractConceptInfoResponse | optional | [ExtractConceptInfoResponse](#gauge.messages.ExtractConceptInfoResponse) |
| formatConceptHeadingRequest | FormatConceptHeadingResponse | optional | [FormatConceptHeadingResponse](#gauge.messages.FormatConceptHeadingResponse) |
| formatConceptHeadingResponse | FormatConceptHeadingResponse | optional | [FormatConceptHeadingResponse](#gauge.messages.FormatConceptHeadingResponse) |

<a name="gauge.messages.ConceptInfo"/>
### ConceptInfo
Details of a Concept

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stepValue | ProtoStepValue | required | The text that defines a concept |
| filepath | string | required | The absolute path to the file that contains the Concept |
| lineNumber | int32 | required | The line number in the file where the concept is defined. |

<a name="gauge.messages.ErrorResponse"/>
### ErrorResponse
A generic failure response

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| error | string | required | Actual error message |

<a name="gauge.messages.ExtractConceptInfoRequest"/>
### ExtractConceptInfoRequest
Request to perform Extract to Concept refactoring
/ The runner does not do the refactoring here, instead it provides inputs enabling the IDE to do refactoring

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| text | string | required | The text blob containing steps that should be refactored to concept. |

<a name="gauge.messages.ExtractConceptInfoResponse"/>
### ExtractConceptInfoResponse
Response to ExtractConceptInfoRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| isValid | bool | required | Flag indicating if the request is valid. |
| hasParam | bool | optional | Flag indicating the presence of parameters |
| conceptHeading | string | optional | The templated concept heading with extracted paramters |
| steps | string | optional | The list of steps that are part of the extracted concept. |
| conceptText | string | optional | The text to invoke the concept from the Spec file/ This contains the value of the parameters to be passed at invocation. |

<a name="gauge.messages.FormatConceptHeadingRequest"/>
### FormatConceptHeadingRequest
Request to change the concept heading, while performing Extract to concept refactoring
/ This is relevant because, the ExtractConceptInfoResponse has a placeholder definition of conceptHeading
/ The user would then give a meaningful name to the concept, and its parameters
/ This is deliberately kept separate from RefactorRequest/RefactorResponse, since those change the files physically.
/ Changing physical files is an expensive affair and requires proper buffer management to support Undo operations.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| newConceptHeading | string | required | The new concept heading |
| oldConceptHeading | string | required | The current Concept Heading that is to be replaced |
| oldConceptText | string | required | The full text of the concept, including the steps |

<a name="gauge.messages.FormatConceptHeadingResponse"/>
### FormatConceptHeadingResponse
Response to FormatConceptHeadingRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| newConceptText | string | required | The new concept text, with revised concept heading. |

<a name="gauge.messages.GetAllConceptsRequest"/>
### GetAllConceptsRequest
Request to get all Concepts in the project

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.GetAllConceptsResponse"/>
### GetAllConceptsResponse
Response to GetAllConceptsResponse

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| concepts | ConceptInfo | repeated | Holds a collection of Concepts that are defined in the project. |

<a name="gauge.messages.GetAllSpecsRequest"/>
### GetAllSpecsRequest
Request to get all Specs in the project

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.GetAllSpecsResponse"/>
### GetAllSpecsResponse
Response to GetAllSpecsRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| specs | ProtoSpec | repeated | Holds a collection of Specs that are defined in the project. |

<a name="gauge.messages.GetAllStepsRequest"/>
### GetAllStepsRequest
Request to get all Steps in the project

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.GetAllStepsResponse"/>
### GetAllStepsResponse
Response to GetAllStepsRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| allSteps | ProtoStepValue | repeated | Holds a collection of Steps that are defined in the project. |

<a name="gauge.messages.GetInstallationRootRequest"/>
### GetInstallationRootRequest
Request to get the Root Directory of the Gauge installation

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.GetInstallationRootResponse"/>
### GetInstallationRootResponse
Response of GetInstallationRootRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| installationRoot | string | required | Holds the absolute path of the Gauge installation directory |

<a name="gauge.messages.GetLanguagePluginLibPathRequest"/>
### GetLanguagePluginLibPathRequest
Request to get the location of language plugin's Lib directory

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| language | string | required | The language to locate the lib directory for. |

<a name="gauge.messages.GetLanguagePluginLibPathResponse"/>
### GetLanguagePluginLibPathResponse
Response to GetLanguagePluginLibPathRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| path | string | required | Absolute apth to the Lib directory of the language. |

<a name="gauge.messages.GetProjectRootRequest"/>
### GetProjectRootRequest
Request to get the Root Directory of the project

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.GetProjectRootResponse"/>
### GetProjectRootResponse
Response of GetProjectRootRequest.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| projectRoot | string | required | Holds the absolute path of the Project Root directory. |

<a name="gauge.messages.GetStepValueRequest"/>
### GetStepValueRequest
Request to get a Step Value.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stepText | string | required | The text of the Step. |
| hasInlineTable | bool | optional | Flag to indicate if the Step has an inline table. |

<a name="gauge.messages.GetStepValueResponse"/>
### GetStepValueResponse
Response to GetStepValueRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stepValue | ProtoStepValue | required | The Step corresponding to the request provided. |

<a name="gauge.messages.PerformRefactoringRequest"/>
### PerformRefactoringRequest
Request to perform a Refactor

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| oldStep | string | required | Step to refactor |
| newStep | string | required | Change to be made |

<a name="gauge.messages.PerformRefactoringResponse"/>
### PerformRefactoringResponse
Response to PerformRefactoringRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| success | bool | required | Flag indicating Success |
| errors | string | repeated | Error message if the refactoring was unsuccessful. |
| filesChanged | string | repeated | Collection of files that were changed as part of the Refactoring. |


<a name="gauge.messages.APIMessage.APIMessageType"/>
### APIMessage.APIMessageType


| Name | Number | Description |
|:----:|:------:|:------------------------------------------------------:|
| GetProjectRootRequest | 1 |  |
| GetProjectRootResponse | 2 |  |
| GetInstallationRootRequest | 3 |  |
| GetInstallationRootResponse | 4 |  |
| GetAllStepsRequest | 5 |  |
| GetAllStepResponse | 6 |  |
| GetAllSpecsRequest | 7 |  |
| GetAllSpecsResponse | 8 |  |
| GetStepValueRequest | 9 |  |
| GetStepValueResponse | 10 |  |
| GetLanguagePluginLibPathRequest | 11 |  |
| GetLanguagePluginLibPathResponse | 12 |  |
| ErrorResponse | 13 |  |
| GetAllConceptsRequest | 14 |  |
| GetAllConceptsResponse | 15 |  |
| PerformRefactoringRequest | 16 |  |
| PerformRefactoringResponse | 17 |  |
| ExtractConceptInfoRequest | 18 |  |
| ExtractConceptInfoResponse | 19 |  |
| FormatConceptHeadingRequest | 20 |  |
| FormatConceptHeadingResponse | 21 |  |

<a name="messages.proto"/>
<p align="right"><a href="#top">Top</a></p>

## messages.proto

<a name="gauge.messages.ExecuteStepRequest"/>
### ExecuteStepRequest
Request sent ot the runner to Execute a Step

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| actualStepText | string | required | Contains the actual text of the Step being executed. / This contains the parameters as defined in the Spec. |
| parsedStepText | string | required | Contains the parsed text of the Step being executed. / The paramters are replaced with placeholders. |
| scenarioFailing | bool | optional | Flag to indicate if the execution of the Scenario, containing the current Step, failed. |
| parameters | Parameter | repeated | Collection of parameters applicable to the current Step. |

<a name="gauge.messages.ExecutionEndingRequest"/>
### ExecutionEndingRequest
Sent at end of Suite Execution. Tells the runner to execute `after_suite` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.ExecutionInfo"/>
### ExecutionInfo
Contains details of the execution. 
/ Depending on the context (Step, Scenario, Spec or Suite), the respective fields are set.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentSpec | SpecInfo | optional | Holds the information of the current Spec. Valid in context of Spec execution. |
| currentScenario | ScenarioInfo | optional | Holds the information of the current Scenario. Valid in context of Scenario execution. |
| currentStep | StepInfo | optional | Holds the information of the current Step. Valid in context of Step execution. |
| stacktrace | string | optional | Stacktrace of the execution. Valid only if there is an error in execution. |

<a name="gauge.messages.ExecutionStartingRequest"/>
### ExecutionStartingRequest
Sent at start of Suite Execution. Tells the runner to execute `before_suite` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.ExecutionStatusResponse"/>
### ExecutionStatusResponse
Sends to any request which needs a execution status as response
/ usually step execution, hooks etc will return this

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| executionResult | ProtoExecutionResult | required |  |

<a name="gauge.messages.KillProcessRequest"/>
### KillProcessRequest
Default request. Tells the runner to shutdown.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.Message"/>
### Message
This is the message which gets transferred all the time
/ with proper message type set
/ One of the Request/Response fields will have value, depending on the MessageType set.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| messageType | Message.MessageType | required |  |
| messageId | int64 | required | A unique id to represent this message. A response to the message should copy over this value./ This is used to synchronize messages &amp; responses |
| executionStartingRequest | ExecutionStartingRequest | optional | [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest) |
| specExecutionStartingRequest | SpecExecutionStartingRequest | optional | [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest) |
| specExecutionEndingRequest | SpecExecutionEndingRequest | optional | [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest) |
| scenarioExecutionStartingRequest | ScenarioExecutionStartingRequest | optional | [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest) |
| scenarioExecutionEndingRequest | ScenarioExecutionEndingRequest | optional | [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest) |
| stepExecutionStartingRequest | StepExecutionStartingRequest | optional | [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest) |
| stepExecutionEndingRequest | StepExecutionEndingRequest | optional | [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest) |
| executeStepRequest | ExecuteStepRequest | optional | [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest) |
| executionEndingRequest | ExecutionEndingRequest | optional | [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest) |
| stepValidateRequest | StepValidateRequest | optional | [StepValidateRequest](#gauge.messages.StepValidateRequest) |
| stepValidateResponse | StepValidateResponse | optional | [StepValidateResponse](#gauge.messages.StepValidateResponse) |
| executionStatusResponse | ExecutionStatusResponse | optional | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) |
| stepNamesRequest | StepNamesRequest | optional | [StepNamesRequest](#gauge.messages.StepNamesRequest) |
| stepNamesResponse | StepNamesResponse | optional | [StepNamesResponse](#gauge.messages.StepNamesResponse) |
| suiteExecutionResult | SuiteExecutionResult | optional | [SuiteExecutionResult ](#gauge.messages.SuiteExecutionResult ) |
| killProcessRequest | KillProcessRequest | optional | [KillProcessRequest](#gauge.messages.KillProcessRequest) |
| scenarioDataStoreInitRequest | ScenarioDataStoreInitRequest | optional | [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest) |
| specDataStoreInitRequest | SpecDataStoreInitRequest | optional | [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest) |
| suiteDataStoreInitRequest | SuiteDataStoreInitRequest | optional | [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest) |
| stepNameRequest | StepNameRequest | optional | [StepNameRequest](#gauge.messages.StepNameRequest) |
| stepNameResponse | StepNameResponse | optional | [StepNameResponse](#gauge.messages.StepNameResponse) |
| refactorRequest | RefactorRequest | optional | [RefactorRequest](#gauge.messages.RefactorRequest) |
| refactorResponse | RefactorResponse | optional | [RefactorResponse](#gauge.messages.RefactorResponse) |

<a name="gauge.messages.ParameterPosition"/>
### ParameterPosition
Holds the new and old positions of a parameter.
/ Used when refactoring a Step.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| oldPosition | int32 | required |  |
| newPosition | int32 | required |  |

<a name="gauge.messages.RefactorRequest"/>
### RefactorRequest
Tells the runner to refactor the specified Step.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| oldStepValue | ProtoStepValue | required | Old value, used to lookup Step to refactor |
| newStepValue | ProtoStepValue | required | New value, the to-be value of Step being refactored. |
| paramPositions | ParameterPosition | repeated | Holds parameter positions of all parameters. Contains old and new parameter positions. |

<a name="gauge.messages.RefactorResponse"/>
### RefactorResponse
Response of a RefactorRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| success | bool | required | Flag indicating the success of Refactor operation. |
| error | string | optional | Error message, valid only if Refactor wasn't successful |
| filesChanged | string | repeated | List of files that were affected because of the refactoring. |

<a name="gauge.messages.ScenarioDataStoreInitRequest"/>
### ScenarioDataStoreInitRequest
Request runner to initialize Scenario DataStore
/ Scenario Datastore is reset after every Scenario execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.ScenarioExecutionEndingRequest"/>
### ScenarioExecutionEndingRequest
Sent at end of Scenario Execution. Tells the runner to execute `after_scenario` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.ScenarioExecutionStartingRequest"/>
### ScenarioExecutionStartingRequest
Sent at start of Scenario Execution. Tells the runner to execute `before_scenario` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.ScenarioInfo"/>
### ScenarioInfo
Contains details of the Scenario execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| name | string | required | Name of the current Scenario being executed. |
| isFailed | bool | required | Flag to indicate if the current Scenario execution failed. |
| tags | string | repeated | Tags relevant to the current Scenario execution. |

<a name="gauge.messages.SpecDataStoreInitRequest"/>
### SpecDataStoreInitRequest
Request runner to initialize Spec DataStore
/ Spec Datastore is reset after every Spec execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.SpecExecutionEndingRequest"/>
### SpecExecutionEndingRequest
Sent at end of Spec Execution. Tells the runner to execute `after_spec` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.SpecExecutionStartingRequest"/>
### SpecExecutionStartingRequest
Sent at start of Spec Execution. Tells the runner to execute `before_spec` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.SpecInfo"/>
### SpecInfo
Contains details of the Spec execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| name | string | required | Name of the current Spec being executed. |
| fileName | string | required | Full File path containing the current Spec being executed. |
| isFailed | bool | required | Flag to indicate if the current Spec execution failed. |
| tags | string | repeated | Tags relevant to the current Spec execution. |

<a name="gauge.messages.StepExecutionEndingRequest"/>
### StepExecutionEndingRequest
Sent at end of Step Execution. Tells the runner to execute `after_step` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.StepExecutionStartingRequest"/>
### StepExecutionStartingRequest
Sent at start of Step Execution. Tells the runner to execute `before_step` hook.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| currentExecutionInfo | ExecutionInfo | optional |  |

<a name="gauge.messages.StepInfo"/>
### StepInfo
Contains details of the Step execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| step | ExecuteStepRequest | required | The current request to execute Step |
| isFailed | bool | required | Flag to indicate if the current Step execution failed. |

<a name="gauge.messages.StepNameRequest"/>
### StepNameRequest
Request for details on a Single Step.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stepValue | string | required | Step text to lookup the Step. / This is the parsed step value, i.e. with placeholders for parameters. |

<a name="gauge.messages.StepNameResponse"/>
### StepNameResponse
Response to StepNameRequest.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| isStepPresent | bool | required | Flag indicating if there is a match for the given Step Text. |
| stepName | string | repeated | The Step name of the given step. |
| hasAlias | bool | required | Flag indicating if the given Step is an alias. |

<a name="gauge.messages.StepNamesRequest"/>
### StepNamesRequest
Requests Gauge to give all Step Names.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.StepNamesResponse"/>
### StepNamesResponse
Response to StepNamesRequest

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| steps | string | repeated | Collection of strings corresponding to Step texts. |

<a name="gauge.messages.StepValidateRequest"/>
### StepValidateRequest
Request sent ot the runner to check if given Step is valid. 
/ The runner should check if there is an implementation defined for the given Step Text.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stepText | string | required | The text is used to lookup Step implementation |
| numberOfParameters | int32 | required | The number of paramters in the Step |

<a name="gauge.messages.StepValidateResponse"/>
### StepValidateResponse
Response of StepValidateRequest.
/ The runner tells the caller if the Request was valid, 
/ i.e. an implementation exists for given Step text.
/ Returns an error message if it is an error response.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| isValid | bool | required |  |
| errorMessage | string | optional |  |

<a name="gauge.messages.SuiteDataStoreInitRequest"/>
### SuiteDataStoreInitRequest
Request runner to initialize Suite DataStore
/ Suite Datastore is reset after every Suite execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|

<a name="gauge.messages.SuiteExecutionResult"/>
### SuiteExecutionResult
Result of the Suite Execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| suiteResult | ProtoSuiteResult | required |  |


<a name="gauge.messages.Message.MessageType"/>
### Message.MessageType


| Name | Number | Description |
|:----:|:------:|:------------------------------------------------------:|
| ExecutionStarting | 0 |  |
| SpecExecutionStarting | 1 |  |
| SpecExecutionEnding | 2 |  |
| ScenarioExecutionStarting | 3 |  |
| ScenarioExecutionEnding | 4 |  |
| StepExecutionStarting | 5 |  |
| StepExecutionEnding | 6 |  |
| ExecuteStep | 7 |  |
| ExecutionEnding | 8 |  |
| StepValidateRequest | 9 |  |
| StepValidateResponse | 10 |  |
| ExecutionStatusResponse | 11 |  |
| StepNamesRequest | 12 |  |
| StepNamesResponse | 13 |  |
| KillProcessRequest | 14 |  |
| SuiteExecutionResult | 15 |  |
| ScenarioDataStoreInit | 16 |  |
| SpecDataStoreInit | 17 |  |
| SuiteDataStoreInit | 18 |  |
| StepNameRequest | 19 |  |
| StepNameResponse | 20 |  |
| RefactorRequest | 21 |  |
| RefactorResponse | 22 |  |

<a name="spec.proto"/>
<p align="right"><a href="#top">Top</a></p>

## spec.proto

<a name="gauge.messages.Fragment"/>
### Fragment
A proto object representing Fragment.
/ Fragments, put together make up A Step

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| fragmentType | Fragment.FragmentType | required | Type of Fragment, valid values are Text, Parameter |
| text | string | optional | Text part of the Fragment, valid only if FragmentType=Text |
| parameter | Parameter | optional | Parameter part of the Fragment, valid only if FragmentType=Parameter |

<a name="gauge.messages.Parameter"/>
### Parameter
A proto object representing Fragment.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| parameterType | Parameter.ParameterType | required | Type of the Parameter. Valid values: Static, Dynamic, Special_String, Special_Table, Table |
| value | string | optional | Holds the value of the parameter |
| name | string | optional | Holds the name of the parameter, used as Key to lookup the value. |
| table | ProtoTable | optional | Holds the table value, if parameterType=Table or Special_Table |

<a name="gauge.messages.ProtoComment"/>
### ProtoComment
A proto object representing Comment.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| text | string | required | Text representing the Comment. |

<a name="gauge.messages.ProtoConcept"/>
### ProtoConcept
Concept is a type of step, that can have multiple Steps.
/ But from a caller's perspective, it is still used as any other Step
/ A proto object representing a Concept

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| conceptStep | ProtoStep | required | Represents the Step value of a Concept. |
| steps | ProtoItem | repeated | Collection of Steps in the given concepts. |
| conceptExecutionResult | ProtoStepExecutionResult | optional | Holds the execution result. |

<a name="gauge.messages.ProtoExecutionResult"/>
### ProtoExecutionResult
A proto object representing the result of an execution

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| failed | bool | required | Flag to indicate failure |
| recoverableError | bool | optional | Flag to indicate if the error is recoverable from. |
| errorMessage | string | optional | The actual error message. |
| stackTrace | string | optional | Stacktrace of the error |
| screenShot | bytes | optional | Byte array containing screenshot taken at the time of failure. |
| executionTime | int64 | required | Holds the time taken for executing this scenario. |

<a name="gauge.messages.ProtoHookFailure"/>
### ProtoHookFailure
A proto object representing a pre-hook failure.
/ Used to hold failure information for before_suite, before_spec, before_scenario and before_spec hooks.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stackTrace | string | required | Stacktrace from the failure |
| errorMessage | string | required | Error message from the failure |
| screenShot | bytes | optional | Byte array holding the screenshot taken at the time of failure. |

<a name="gauge.messages.ProtoItem"/>
### ProtoItem
Container for all valid Items under a Specification.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| itemType | ProtoItem.ItemType | required | Itemtype of the current ProtoItem |
| step | ProtoStep | optional | Holds the Step definition. Valid only if ItemType = Step |
| concept | ProtoConcept | optional | Holds the Concept definition. Valid only if ItemType = Concept |
| scenario | ProtoScenario | optional | Holds the Scenario definition. Valid only if ItemType = Scenario |
| tableDrivenScenario | ProtoTableDrivenScenario | optional | Holds the TableDrivenScenario definition. Valid only if ItemType = TableDrivenScenario |
| comment | ProtoComment | optional | Holds the Comment definition. Valid only if ItemType = Comment |
| table | ProtoTable | optional | Holds the Table definition. Valid only if ItemType = Table |
| tags | ProtoTags | optional | Holds the Tags definition. Valid only if ItemType = Tags |

<a name="gauge.messages.ProtoScenario"/>
### ProtoScenario
A proto object representing a Scenario

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| scenarioHeading | string | required | Heading of the given Scenario |
| failed | bool | required | Flag to indicate if the Scenario execution failed |
| contexts | ProtoItem | repeated | Collection of Context steps. The Context steps are executed before every run. |
| scenarioItems | ProtoItem | repeated | Collection of Items under a scenario. These could be Steps, Comments, Tags, TableDrivenScenarios or Tables |
| preHookFailure | ProtoHookFailure | optional | Contains a 'before' hook failure message. This happens when the `before_scenario` hook has an error. |
| postHookFailure | ProtoHookFailure | optional | Contains a 'after' hook failure message. This happens when the `after_scenario` hook has an error. |
| tags | string | repeated | Contains a list of tags that are defined at the specification level. Scenario tags are not present here. |
| executionTime | int64 | optional | Holds the time taken for executing this scenario. |

<a name="gauge.messages.ProtoSpec"/>
### ProtoSpec
A proto object representing a Specification
/ A specification can contain Scenarios or Steps, besides Comments

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| specHeading | string | required | Heading describing the Specification |
| items | ProtoItem | repeated | A collection of items that come under this step |
| isTableDriven | bool | required | Flag indicating if this is a Table Driven Specification. The table is defined in the context, this is different from using a table parameter. |
| preHookFailure | ProtoHookFailure | optional | Contains a 'before' hook failure message. This happens when the `before_spec` hook has an error. |
| postHookFailure | ProtoHookFailure | optional | Contains a 'before' hook failure message. This happens when the `after_hook` hook has an error. |
| fileName | string | required | Contains the filename for that holds this specification. |
| tags | string | repeated | Contains a list of tags that are defined at the specification level. Scenario tags are not present here. |

<a name="gauge.messages.ProtoSpecResult"/>
### ProtoSpecResult
A proto object representing the result of Spec execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| protoSpec | ProtoSpec | required | Represents the corresponding Specification |
| scenarioCount | int32 | required | Holds the number of Scenarios executed |
| scenarioFailedCount | int32 | required | Holds the number of Scenarios failed |
| failed | bool | required | Flag to indicate failure |
| failedDataTableRows | int32 | repeated | Holds the row numbers, which caused the execution to fail. |
| executionTime | int64 | optional | Holds the time taken for executing the spec. |

<a name="gauge.messages.ProtoStep"/>
### ProtoStep
A proto object representing a Step

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| actualText | string | required | Holds the raw text of the Step as defined in the spec file. This contains the actual parameter values. |
| parsedText | string | required | Contains the parsed text of the Step. This will have placeholders for the parameters. |
| fragments | Fragment | repeated | Collection of a list of fragments for a Step. A fragment could be either text or parameter. |
| stepExecutionResult | ProtoStepExecutionResult | optional | Holds the result from the execution. |

<a name="gauge.messages.ProtoStepExecutionResult"/>
### ProtoStepExecutionResult
A proto object representing Step Execution result

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| executionResult | ProtoExecutionResult | required | The actual result of the execution |
| preHookFailure | ProtoHookFailure | optional | Contains a 'before' hook failure message. This happens when the `before_step` hook has an error. |
| postHookFailure | ProtoHookFailure | optional | Contains a 'after' hook failure message. This happens when the `after_step` hook has an error. |

<a name="gauge.messages.ProtoStepValue"/>
### ProtoStepValue
A proto object representing a Step value.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| stepValue | string | required | The actual string value describing he Step |
| parameterizedStepValue | string | required | The parameterized string value describing he Step. The parameters are replaced with placeholders. |
| parameters | string | repeated | A collection of strings representing the parameters. |

<a name="gauge.messages.ProtoSuiteResult"/>
### ProtoSuiteResult
A proto object representing the result of entire Suite execution.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| specResults | ProtoSpecResult | repeated | Contains the result from the execution |
| preHookFailure | ProtoHookFailure | optional | Contains a 'before' hook failure message. This happens when the `before_suite` hook has an error |
| postHookFailure | ProtoHookFailure | optional | Contains a 'after' hook failure message. This happens when the `after_suite` hook has an error |
| failed | bool | required | Flag to indicate failure |
| specsFailedCount | int32 | required | Holds the count of number of Specifications that failed. |
| executionTime | int64 | optional | Holds the time taken for executing the whole suite. |
| successRate | float | required | Holds a metric indicating the success rate of the execution. |

<a name="gauge.messages.ProtoTable"/>
### ProtoTable
A proto object representing Table.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| headers | ProtoTableRow | required | Contains the Headers for the table |
| rows | ProtoTableRow | repeated | Contains the Rows for the table |

<a name="gauge.messages.ProtoTableDrivenScenario"/>
### ProtoTableDrivenScenario
A proto object representing a TableDrivenScenario

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| scenarios | ProtoScenario | repeated | Holds the Underlying scenario that is executed for every row in the table. |

<a name="gauge.messages.ProtoTableRow"/>
### ProtoTableRow
A proto object representing Table.

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| cells | string | repeated | Represents the cells of a given table |

<a name="gauge.messages.ProtoTags"/>
### ProtoTags
A proto object representing Tags

| Field | Type | Label | Description |
|:-----:|:----:|:-----:|:------------------------------------------------------:|
| tags | string | repeated | A collection of Tags |


<a name="gauge.messages.Fragment.FragmentType"/>
### Fragment.FragmentType
Enum representing the types of Fragment

| Name | Number | Description |
|:----:|:------:|:------------------------------------------------------:|
| Text | 1 | Fragment is a Text part |
| Parameter | 2 | Fragment is a Parameter part |

<a name="gauge.messages.Parameter.ParameterType"/>
### Parameter.ParameterType
Enum representing types of Parameter.

| Name | Number | Description |
|:----:|:------:|:------------------------------------------------------:|
| Static | 1 |  |
| Dynamic | 2 |  |
| Special_String | 3 |  |
| Special_Table | 4 |  |
| Table | 5 |  |

<a name="gauge.messages.ProtoItem.ItemType"/>
### ProtoItem.ItemType
Enumerates various item types that the proto item can contain. Valid types are: Step, Comment, Concept, Scenario, TableDrivenScenario, Table, Tags

| Name | Number | Description |
|:----:|:------:|:------------------------------------------------------:|
| Step | 1 |  |
| Comment | 2 |  |
| Concept | 3 |  |
| Scenario | 4 |  |
| TableDrivenScenario | 5 |  |
| Table | 6 |  |
| Tags | 7 |  |


<a name="scalar-value-types"/>
## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| double |  | double | double | float |
| float |  | float | float | float |
| int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| sfixed32 | Always four bytes. | int32 | int | int |
| sfixed64 | Always eight bytes. | int64 | long | int/long |
| bool |  | bool | boolean | boolean |
| string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |
