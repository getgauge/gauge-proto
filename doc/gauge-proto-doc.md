# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api.proto)
    - [APIMessage](#gauge.messages.APIMessage)
    - [ConceptInfo](#gauge.messages.ConceptInfo)
    - [ErrorResponse](#gauge.messages.ErrorResponse)
    - [ExtractConceptRequest](#gauge.messages.ExtractConceptRequest)
    - [ExtractConceptResponse](#gauge.messages.ExtractConceptResponse)
    - [FormatSpecsRequest](#gauge.messages.FormatSpecsRequest)
    - [FormatSpecsResponse](#gauge.messages.FormatSpecsResponse)
    - [GetAllConceptsRequest](#gauge.messages.GetAllConceptsRequest)
    - [GetAllConceptsResponse](#gauge.messages.GetAllConceptsResponse)
    - [GetAllStepsRequest](#gauge.messages.GetAllStepsRequest)
    - [GetAllStepsResponse](#gauge.messages.GetAllStepsResponse)
    - [GetInstallationRootRequest](#gauge.messages.GetInstallationRootRequest)
    - [GetInstallationRootResponse](#gauge.messages.GetInstallationRootResponse)
    - [GetLanguagePluginLibPathRequest](#gauge.messages.GetLanguagePluginLibPathRequest)
    - [GetLanguagePluginLibPathResponse](#gauge.messages.GetLanguagePluginLibPathResponse)
    - [GetProjectRootRequest](#gauge.messages.GetProjectRootRequest)
    - [GetProjectRootResponse](#gauge.messages.GetProjectRootResponse)
    - [GetStepValueRequest](#gauge.messages.GetStepValueRequest)
    - [GetStepValueResponse](#gauge.messages.GetStepValueResponse)
    - [PerformRefactoringRequest](#gauge.messages.PerformRefactoringRequest)
    - [PerformRefactoringResponse](#gauge.messages.PerformRefactoringResponse)
    - [SpecsRequest](#gauge.messages.SpecsRequest)
    - [SpecsResponse](#gauge.messages.SpecsResponse)
    - [SpecsResponse.SpecDetail](#gauge.messages.SpecsResponse.SpecDetail)
    - [UnsupportedApiMessageResponse](#gauge.messages.UnsupportedApiMessageResponse)
    - [step](#gauge.messages.step)
    - [textInfo](#gauge.messages.textInfo)
  
    - [APIMessage.APIMessageType](#gauge.messages.APIMessage.APIMessageType)
  
- [lsp.proto](#lsp.proto)
    - [lspService](#gauge.messages.lspService)
  
- [messages.proto](#messages.proto)
    - [CacheFileRequest](#gauge.messages.CacheFileRequest)
    - [Empty](#gauge.messages.Empty)
    - [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest)
    - [ExecutionArg](#gauge.messages.ExecutionArg)
    - [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest)
    - [ExecutionInfo](#gauge.messages.ExecutionInfo)
    - [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest)
    - [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse)
    - [FileChanges](#gauge.messages.FileChanges)
    - [FileDiff](#gauge.messages.FileDiff)
    - [ImplementationFileGlobPatternRequest](#gauge.messages.ImplementationFileGlobPatternRequest)
    - [ImplementationFileGlobPatternResponse](#gauge.messages.ImplementationFileGlobPatternResponse)
    - [ImplementationFileListRequest](#gauge.messages.ImplementationFileListRequest)
    - [ImplementationFileListResponse](#gauge.messages.ImplementationFileListResponse)
    - [KeepAlive](#gauge.messages.KeepAlive)
    - [KillProcessRequest](#gauge.messages.KillProcessRequest)
    - [Message](#gauge.messages.Message)
    - [ParameterPosition](#gauge.messages.ParameterPosition)
    - [RefactorRequest](#gauge.messages.RefactorRequest)
    - [RefactorResponse](#gauge.messages.RefactorResponse)
    - [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest)
    - [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest)
    - [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest)
    - [ScenarioInfo](#gauge.messages.ScenarioInfo)
    - [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest)
    - [SpecDetails](#gauge.messages.SpecDetails)
    - [SpecDetails.SpecDetail](#gauge.messages.SpecDetails.SpecDetail)
    - [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest)
    - [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest)
    - [SpecInfo](#gauge.messages.SpecInfo)
    - [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest)
    - [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest)
    - [StepInfo](#gauge.messages.StepInfo)
    - [StepNameRequest](#gauge.messages.StepNameRequest)
    - [StepNameResponse](#gauge.messages.StepNameResponse)
    - [StepNamesRequest](#gauge.messages.StepNamesRequest)
    - [StepNamesResponse](#gauge.messages.StepNamesResponse)
    - [StepPositionsRequest](#gauge.messages.StepPositionsRequest)
    - [StepPositionsResponse](#gauge.messages.StepPositionsResponse)
    - [StepPositionsResponse.StepPosition](#gauge.messages.StepPositionsResponse.StepPosition)
    - [StepValidateRequest](#gauge.messages.StepValidateRequest)
    - [StepValidateResponse](#gauge.messages.StepValidateResponse)
    - [StubImplementationCodeRequest](#gauge.messages.StubImplementationCodeRequest)
    - [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest)
    - [SuiteExecutionResult](#gauge.messages.SuiteExecutionResult)
    - [SuiteExecutionResultItem](#gauge.messages.SuiteExecutionResultItem)
    - [TextDiff](#gauge.messages.TextDiff)
    - [UnsupportedMessageResponse](#gauge.messages.UnsupportedMessageResponse)
  
    - [CacheFileRequest.FileStatus](#gauge.messages.CacheFileRequest.FileStatus)
    - [Message.MessageType](#gauge.messages.Message.MessageType)
    - [StepValidateResponse.ErrorType](#gauge.messages.StepValidateResponse.ErrorType)
  
- [services.proto](#services.proto)
    - [Documenter](#gauge.messages.Documenter)
    - [Reporter](#gauge.messages.Reporter)
    - [Runner](#gauge.messages.Runner)
  
- [spec.proto](#spec.proto)
    - [Error](#gauge.messages.Error)
    - [Fragment](#gauge.messages.Fragment)
    - [Parameter](#gauge.messages.Parameter)
    - [ProtoComment](#gauge.messages.ProtoComment)
    - [ProtoConcept](#gauge.messages.ProtoConcept)
    - [ProtoExecutionResult](#gauge.messages.ProtoExecutionResult)
    - [ProtoHookFailure](#gauge.messages.ProtoHookFailure)
    - [ProtoItem](#gauge.messages.ProtoItem)
    - [ProtoScenario](#gauge.messages.ProtoScenario)
    - [ProtoScenarioResult](#gauge.messages.ProtoScenarioResult)
    - [ProtoSpec](#gauge.messages.ProtoSpec)
    - [ProtoSpecResult](#gauge.messages.ProtoSpecResult)
    - [ProtoStep](#gauge.messages.ProtoStep)
    - [ProtoStepExecutionResult](#gauge.messages.ProtoStepExecutionResult)
    - [ProtoStepResult](#gauge.messages.ProtoStepResult)
    - [ProtoStepValue](#gauge.messages.ProtoStepValue)
    - [ProtoSuiteResult](#gauge.messages.ProtoSuiteResult)
    - [ProtoTable](#gauge.messages.ProtoTable)
    - [ProtoTableDrivenScenario](#gauge.messages.ProtoTableDrivenScenario)
    - [ProtoTableRow](#gauge.messages.ProtoTableRow)
    - [ProtoTags](#gauge.messages.ProtoTags)
    - [Span](#gauge.messages.Span)
  
    - [Error.ErrorType](#gauge.messages.Error.ErrorType)
    - [ExecutionStatus](#gauge.messages.ExecutionStatus)
    - [Fragment.FragmentType](#gauge.messages.Fragment.FragmentType)
    - [Parameter.ParameterType](#gauge.messages.Parameter.ParameterType)
    - [ProtoExecutionResult.ErrorType](#gauge.messages.ProtoExecutionResult.ErrorType)
    - [ProtoItem.ItemType](#gauge.messages.ProtoItem.ItemType)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto



<a name="gauge.messages.APIMessage"></a>

### APIMessage
A generic message composing of all possible operations.
/ One of the Request/Response fields will have value, depending on the MessageType set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messageType | [APIMessage.APIMessageType](#gauge.messages.APIMessage.APIMessageType) |  | Type of API call being made |
| messageId | [int64](#int64) |  | A unique id to represent this message. A response to the message should copy over this value. / This is used to synchronize messages &amp; responses |
| projectRootRequest | [GetProjectRootRequest](#gauge.messages.GetProjectRootRequest) |  | [GetProjectRootRequest](#gauge.messages.GetProjectRootRequest) |
| projectRootResponse | [GetProjectRootResponse](#gauge.messages.GetProjectRootResponse) |  | [GetProjectRootResponse](#gauge.messages.GetProjectRootResponse) |
| installationRootRequest | [GetInstallationRootRequest](#gauge.messages.GetInstallationRootRequest) |  | [GetInstallationRootRequest](#gauge.messages.GetInstallationRootRequest) |
| installationRootResponse | [GetInstallationRootResponse](#gauge.messages.GetInstallationRootResponse) |  | [GetInstallationRootResponse](#gauge.messages.GetInstallationRootResponse) |
| allStepsRequest | [GetAllStepsRequest](#gauge.messages.GetAllStepsRequest) |  | [GetAllStepsRequest](#gauge.messages.GetAllStepsRequest) |
| allStepsResponse | [GetAllStepsResponse](#gauge.messages.GetAllStepsResponse) |  | [GetAllStepsResponse](#gauge.messages.GetAllStepsResponse) |
| specsRequest | [SpecsRequest](#gauge.messages.SpecsRequest) |  | [GetAllSpecsRequest](#gauge.messages.GetAllSpecsRequest) |
| specsResponse | [SpecsResponse](#gauge.messages.SpecsResponse) |  | [GetAllSpecsResponse](#gauge.messages.GetAllSpecsResponse) |
| stepValueRequest | [GetStepValueRequest](#gauge.messages.GetStepValueRequest) |  | [GetStepValueRequest](#gauge.messages.GetStepValueRequest) |
| stepValueResponse | [GetStepValueResponse](#gauge.messages.GetStepValueResponse) |  | [GetStepValueResponse](#gauge.messages.GetStepValueResponse) |
| libPathRequest | [GetLanguagePluginLibPathRequest](#gauge.messages.GetLanguagePluginLibPathRequest) |  | [GetLanguagePluginLibPathRequest](#gauge.messages.GetLanguagePluginLibPathRequest) |
| libPathResponse | [GetLanguagePluginLibPathResponse](#gauge.messages.GetLanguagePluginLibPathResponse) |  | [GetLanguagePluginLibPathResponse](#gauge.messages.GetLanguagePluginLibPathResponse) |
| error | [ErrorResponse](#gauge.messages.ErrorResponse) |  | [ErrorResponse](#gauge.messages.ErrorResponse) |
| allConceptsRequest | [GetAllConceptsRequest](#gauge.messages.GetAllConceptsRequest) |  | [GetAllConceptsRequest](#gauge.messages.GetAllConceptsRequest) |
| allConceptsResponse | [GetAllConceptsResponse](#gauge.messages.GetAllConceptsResponse) |  | [GetAllConceptsResponse](#gauge.messages.GetAllConceptsResponse) |
| performRefactoringRequest | [PerformRefactoringRequest](#gauge.messages.PerformRefactoringRequest) |  | [PerformRefactoringRequest](#gauge.messages.PerformRefactoringRequest) |
| performRefactoringResponse | [PerformRefactoringResponse](#gauge.messages.PerformRefactoringResponse) |  | [PerformRefactoringResponse](#gauge.messages.PerformRefactoringResponse) |
| extractConceptRequest | [ExtractConceptRequest](#gauge.messages.ExtractConceptRequest) |  | [ExtractConceptRequest](#gauge.messages.ExtractConceptRequest) |
| extractConceptResponse | [ExtractConceptResponse](#gauge.messages.ExtractConceptResponse) |  | [ExtractConceptResponse](#gauge.messages.ExtractConceptResponse) |
| formatSpecsRequest | [FormatSpecsRequest](#gauge.messages.FormatSpecsRequest) |  | [FormatSpecsRequest] (#gauge.messages.FormatSpecsRequest) |
| formatSpecsResponse | [FormatSpecsResponse](#gauge.messages.FormatSpecsResponse) |  | [FormatSpecsResponse] (#gauge.messages.FormatSpecsResponse) |
| unsupportedApiMessageResponse | [UnsupportedApiMessageResponse](#gauge.messages.UnsupportedApiMessageResponse) |  | [UnsupportedApiMessageResponse] (#gauge.messages.UnsupportedApiMessageResponse) |






<a name="gauge.messages.ConceptInfo"></a>

### ConceptInfo
Details of a Concept


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepValue | [ProtoStepValue](#gauge.messages.ProtoStepValue) |  | The text that defines a concept |
| filepath | [string](#string) |  | The absolute path to the file that contains the Concept |
| lineNumber | [int32](#int32) |  | The line number in the file where the concept is defined. |






<a name="gauge.messages.ErrorResponse"></a>

### ErrorResponse
A generic failure response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error | [string](#string) |  | Actual error message |






<a name="gauge.messages.ExtractConceptRequest"></a>

### ExtractConceptRequest
Request to perform Extract to Concept refactoring


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| conceptName | [step](#gauge.messages.step) |  | The Concept name given by the user |
| steps | [step](#gauge.messages.step) | repeated | steps to extract |
| changeAcrossProject | [bool](#bool) |  | Flag indicating if refactoring should be done across project |
| conceptFileName | [string](#string) |  | The concept filename in which extracted concept will be added |
| selectedTextInfo | [textInfo](#gauge.messages.textInfo) |  | Info related to selected text, only if changeAcrossProject is false |






<a name="gauge.messages.ExtractConceptResponse"></a>

### ExtractConceptResponse
Response to perform Extract to Concept refactoring


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isSuccess | [bool](#bool) |  | Flag indicating Success |
| error | [string](#string) |  | Error message if the refactoring was unsuccessful. |
| filesChanged | [string](#string) | repeated | Collection of files that were changed as part of the Refactoring. |






<a name="gauge.messages.FormatSpecsRequest"></a>

### FormatSpecsRequest
Request to format spec files


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specs | [string](#string) | repeated | Specs to be formatted |






<a name="gauge.messages.FormatSpecsResponse"></a>

### FormatSpecsResponse
Response on formatting spec files


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| errors | [string](#string) | repeated | Errors occurred on formatting |
| warnings | [string](#string) | repeated | Warnings occurred on formatting |






<a name="gauge.messages.GetAllConceptsRequest"></a>

### GetAllConceptsRequest
Request to get all Concepts in the project






<a name="gauge.messages.GetAllConceptsResponse"></a>

### GetAllConceptsResponse
Response to GetAllConceptsResponse


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| concepts | [ConceptInfo](#gauge.messages.ConceptInfo) | repeated | Holds a collection of Concepts that are defined in the project. |






<a name="gauge.messages.GetAllStepsRequest"></a>

### GetAllStepsRequest
Request to get all Steps in the project






<a name="gauge.messages.GetAllStepsResponse"></a>

### GetAllStepsResponse
Response to GetAllStepsRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| allSteps | [ProtoStepValue](#gauge.messages.ProtoStepValue) | repeated | Holds a collection of Steps that are defined in the project. |






<a name="gauge.messages.GetInstallationRootRequest"></a>

### GetInstallationRootRequest
Request to get the Root Directory of the Gauge installation






<a name="gauge.messages.GetInstallationRootResponse"></a>

### GetInstallationRootResponse
Response of GetInstallationRootRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| installationRoot | [string](#string) |  | Holds the absolute path of the Gauge installation directory |






<a name="gauge.messages.GetLanguagePluginLibPathRequest"></a>

### GetLanguagePluginLibPathRequest
Request to get the location of language plugin&#39;s Lib directory


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| language | [string](#string) |  | The language to locate the lib directory for. |






<a name="gauge.messages.GetLanguagePluginLibPathResponse"></a>

### GetLanguagePluginLibPathResponse
Response to GetLanguagePluginLibPathRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path | [string](#string) |  | Absolute path to the Lib directory of the language. |






<a name="gauge.messages.GetProjectRootRequest"></a>

### GetProjectRootRequest
Request to get the Root Directory of the project






<a name="gauge.messages.GetProjectRootResponse"></a>

### GetProjectRootResponse
Response of GetProjectRootRequest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| projectRoot | [string](#string) |  | Holds the absolute path of the Project Root directory. |






<a name="gauge.messages.GetStepValueRequest"></a>

### GetStepValueRequest
Request to get a Step Value.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepText | [string](#string) |  | The text of the Step. |
| hasInlineTable | [bool](#bool) |  | Flag to indicate if the Step has an inline table. |






<a name="gauge.messages.GetStepValueResponse"></a>

### GetStepValueResponse
Response to GetStepValueRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepValue | [ProtoStepValue](#gauge.messages.ProtoStepValue) |  | The Step corresponding to the request provided. |






<a name="gauge.messages.PerformRefactoringRequest"></a>

### PerformRefactoringRequest
Request to perform a Refactor


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oldStep | [string](#string) |  | Step to refactor |
| newStep | [string](#string) |  | Change to be made |






<a name="gauge.messages.PerformRefactoringResponse"></a>

### PerformRefactoringResponse
Response to PerformRefactoringRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Flag indicating Success |
| errors | [string](#string) | repeated | Error message if the refactoring was unsuccessful. |
| filesChanged | [string](#string) | repeated | Collection of files that were changed as part of the Refactoring. |






<a name="gauge.messages.SpecsRequest"></a>

### SpecsRequest
Request to get all Specs in the project


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specs | [string](#string) | repeated |  |






<a name="gauge.messages.SpecsResponse"></a>

### SpecsResponse
Response to GetAllSpecsRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| details | [SpecsResponse.SpecDetail](#gauge.messages.SpecsResponse.SpecDetail) | repeated | Holds a collection of Spec details. |






<a name="gauge.messages.SpecsResponse.SpecDetail"></a>

### SpecsResponse.SpecDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ProtoSpec](#gauge.messages.ProtoSpec) |  | Holds a collection of Specs that are defined in the project. |
| parseErrors | [Error](#gauge.messages.Error) | repeated | Holds a collection of parse errors present in the above spec. |






<a name="gauge.messages.UnsupportedApiMessageResponse"></a>

### UnsupportedApiMessageResponse
Response when a API message request is not supported.






<a name="gauge.messages.step"></a>

### step



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name of the step |
| table | [string](#string) |  | table present in step as parameter |
| paramTableName | [string](#string) |  | name of table in concept heading, if it comes as a param to concept |






<a name="gauge.messages.textInfo"></a>

### textInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fileName | [string](#string) |  | The filename from where concept is being extracted |
| startingLineNo | [int32](#int32) |  | storing the starting and ending line number of selected text |
| endLineNo | [int32](#int32) |  |  |





 


<a name="gauge.messages.APIMessage.APIMessageType"></a>

### APIMessage.APIMessageType


| Name | Number | Description |
| ---- | ------ | ----------- |
| GetProjectRootRequest | 0 |  |
| GetProjectRootResponse | 1 |  |
| GetInstallationRootRequest | 2 |  |
| GetInstallationRootResponse | 3 |  |
| GetAllStepsRequest | 4 |  |
| GetAllStepResponse | 5 |  |
| SpecsRequest | 6 |  |
| SpecsResponse | 7 |  |
| GetStepValueRequest | 8 |  |
| GetStepValueResponse | 9 |  |
| GetLanguagePluginLibPathRequest | 10 |  |
| GetLanguagePluginLibPathResponse | 11 |  |
| ErrorResponse | 12 |  |
| GetAllConceptsRequest | 13 |  |
| GetAllConceptsResponse | 14 |  |
| PerformRefactoringRequest | 15 |  |
| PerformRefactoringResponse | 16 |  |
| ExtractConceptRequest | 17 |  |
| ExtractConceptResponse | 18 |  |
| FormatSpecsRequest | 19 |  |
| FormatSpecsResponse | 20 |  |
| UnsupportedApiMessageResponse | 21 |  |


 

 

 



<a name="lsp.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## lsp.proto


 

 

 


<a name="gauge.messages.lspService"></a>

### lspService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetStepNames | [StepNamesRequest](#gauge.messages.StepNamesRequest) | [StepNamesResponse](#gauge.messages.StepNamesResponse) |  |
| CacheFile | [CacheFileRequest](#gauge.messages.CacheFileRequest) | [Empty](#gauge.messages.Empty) |  |
| GetStepPositions | [StepPositionsRequest](#gauge.messages.StepPositionsRequest) | [StepPositionsResponse](#gauge.messages.StepPositionsResponse) |  |
| GetImplementationFiles | [Empty](#gauge.messages.Empty) | [ImplementationFileListResponse](#gauge.messages.ImplementationFileListResponse) |  |
| ImplementStub | [StubImplementationCodeRequest](#gauge.messages.StubImplementationCodeRequest) | [FileDiff](#gauge.messages.FileDiff) |  |
| ValidateStep | [StepValidateRequest](#gauge.messages.StepValidateRequest) | [StepValidateResponse](#gauge.messages.StepValidateResponse) |  |
| Refactor | [RefactorRequest](#gauge.messages.RefactorRequest) | [RefactorResponse](#gauge.messages.RefactorResponse) |  |
| GetStepName | [StepNameRequest](#gauge.messages.StepNameRequest) | [StepNameResponse](#gauge.messages.StepNameResponse) |  |
| GetGlobPatterns | [Empty](#gauge.messages.Empty) | [ImplementationFileGlobPatternResponse](#gauge.messages.ImplementationFileGlobPatternResponse) |  |
| KillProcess | [KillProcessRequest](#gauge.messages.KillProcessRequest) | [Empty](#gauge.messages.Empty) |  |

 



<a name="messages.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## messages.proto
The comments are exported to Markdown, hence they may contain markdown syntax and cross-refs.


<a name="gauge.messages.CacheFileRequest"></a>

### CacheFileRequest
Request for caching a file.
/ Gauge sends this request when running in LSP mode,
/ so runner can cache file contents present on the client(an editor).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content | [string](#string) |  | File content of the file to be cached |
| filePath | [string](#string) |  | File path of the file to be cached |
| isClosed | [bool](#bool) |  | Specifies if the file is closed |
| status | [CacheFileRequest.FileStatus](#gauge.messages.CacheFileRequest.FileStatus) |  | Specifies the status of the file |






<a name="gauge.messages.Empty"></a>

### Empty
Empty is a blank response, to be used when there is no return expected.






<a name="gauge.messages.ExecuteStepRequest"></a>

### ExecuteStepRequest
Request sent ot the runner to Execute a Step


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actualStepText | [string](#string) |  | Contains the actual text of the Step being executed. / This contains the parameters as defined in the Spec. |
| parsedStepText | [string](#string) |  | Contains the parsed text of the Step being executed. / The paramters are replaced with placeholders. |
| scenarioFailing | [bool](#bool) |  | Flag to indicate if the execution of the Scenario, containing the current Step, failed. |
| parameters | [Parameter](#gauge.messages.Parameter) | repeated | Collection of parameters applicable to the current Step. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.ExecutionArg"></a>

### ExecutionArg
Contains command line arguments which passed by user during execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| flagName | [string](#string) |  | Holds the flag name passed from command line. |
| flagValue | [string](#string) | repeated | Holds the flag value passed from command line. |






<a name="gauge.messages.ExecutionEndingRequest"></a>

### ExecutionEndingRequest
Sent at end of Suite Execution. Tells the runner to execute `after_suite` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current suite execution info. |
| suiteResult | [ProtoSuiteResult](#gauge.messages.ProtoSuiteResult) |  | Holds the suite result in execution ending. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.ExecutionInfo"></a>

### ExecutionInfo
Contains details of the execution.
/ Depending on the context (Step, Scenario, Spec or Suite), the respective fields are set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentSpec | [SpecInfo](#gauge.messages.SpecInfo) |  | Holds the information of the current Spec. Valid in context of Spec execution. |
| currentScenario | [ScenarioInfo](#gauge.messages.ScenarioInfo) |  | Holds the information of the current Scenario. Valid in context of Scenario execution. |
| currentStep | [StepInfo](#gauge.messages.StepInfo) |  | Holds the information of the current Step. Valid in context of Step execution. |
| stacktrace | [string](#string) |  | Stacktrace of the execution. Valid only if there is an error in execution. |
| projectName | [string](#string) |  | Holds the project name |
| ExecutionArgs | [ExecutionArg](#gauge.messages.ExecutionArg) | repeated | Holds the command line arguments. |
| numberOfExecutionStreams | [int32](#int32) |  | Holds the number of running execution streams. |
| runnerId | [int32](#int32) |  | Holds the runner id for parallel execution. |






<a name="gauge.messages.ExecutionStartingRequest"></a>

### ExecutionStartingRequest
Sent at start of Suite Execution. Tells the runner to execute `before_suite` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current suite execution info. |
| suiteResult | [ProtoSuiteResult](#gauge.messages.ProtoSuiteResult) |  | Holds the suite result in execution starting. / Some fields will not be populated before execution. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.ExecutionStatusResponse"></a>

### ExecutionStatusResponse
Sends to any request which needs a execution status as response
/ usually step execution, hooks etc will return this


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| executionResult | [ProtoExecutionResult](#gauge.messages.ProtoExecutionResult) |  | Holds the suite result after suite execution done. |






<a name="gauge.messages.FileChanges"></a>

### FileChanges
Give all file changes to be made to file system


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fileName | [string](#string) |  |  |
| fileContent | [string](#string) |  |  |
| diffs | [TextDiff](#gauge.messages.TextDiff) | repeated |  |






<a name="gauge.messages.FileDiff"></a>

### FileDiff
Diffs to be applied to a file


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filePath | [string](#string) |  | File Path where the new content needs to be put in |
| textDiffs | [TextDiff](#gauge.messages.TextDiff) | repeated | The diffs which need to be applied to this file |






<a name="gauge.messages.ImplementationFileGlobPatternRequest"></a>

### ImplementationFileGlobPatternRequest
Request for getting Implementation file glob pattern






<a name="gauge.messages.ImplementationFileGlobPatternResponse"></a>

### ImplementationFileGlobPatternResponse
Response for getting Implementation file glob pattern


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| globPatterns | [string](#string) | repeated | List of implementation file glob patterns |






<a name="gauge.messages.ImplementationFileListRequest"></a>

### ImplementationFileListRequest
Request for getting Implementation file list






<a name="gauge.messages.ImplementationFileListResponse"></a>

### ImplementationFileListResponse
Response for getting Implementation file list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| implementationFilePaths | [string](#string) | repeated | List of implementation files |






<a name="gauge.messages.KeepAlive"></a>

### KeepAlive
Tell gauge to reset the kill timer, thus extending the life


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pluginId | [string](#string) |  | ID of the plugin initiating this request |






<a name="gauge.messages.KillProcessRequest"></a>

### KillProcessRequest
Default request. Tells the runner to shutdown.






<a name="gauge.messages.Message"></a>

### Message
This is the message which gets transferred all the time
/ with proper message type set
/ One of the Request/Response fields will have value, depending on the MessageType set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| messageType | [Message.MessageType](#gauge.messages.Message.MessageType) |  |  |
| messageId | [int64](#int64) |  | A unique id to represent this message. A response to the message should copy over this value. / This is used to synchronize messages &amp; responses |
| executionStartingRequest | [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest) |  | [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest) |
| specExecutionStartingRequest | [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest) |  | [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest) |
| specExecutionEndingRequest | [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest) |  | [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest) |
| scenarioExecutionStartingRequest | [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest) |  | [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest) |
| scenarioExecutionEndingRequest | [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest) |  | [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest) |
| stepExecutionStartingRequest | [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest) |  | [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest) |
| stepExecutionEndingRequest | [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest) |  | [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest) |
| executeStepRequest | [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest) |  | [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest) |
| executionEndingRequest | [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest) |  | [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest) |
| stepValidateRequest | [StepValidateRequest](#gauge.messages.StepValidateRequest) |  | [StepValidateRequest](#gauge.messages.StepValidateRequest) |
| stepValidateResponse | [StepValidateResponse](#gauge.messages.StepValidateResponse) |  | [StepValidateResponse](#gauge.messages.StepValidateResponse) |
| executionStatusResponse | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) |  | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) |
| stepNamesRequest | [StepNamesRequest](#gauge.messages.StepNamesRequest) |  | [StepNamesRequest](#gauge.messages.StepNamesRequest) |
| stepNamesResponse | [StepNamesResponse](#gauge.messages.StepNamesResponse) |  | [StepNamesResponse](#gauge.messages.StepNamesResponse) |
| suiteExecutionResult | [SuiteExecutionResult](#gauge.messages.SuiteExecutionResult) |  | [SuiteExecutionResult ](#gauge.messages.SuiteExecutionResult ) |
| killProcessRequest | [KillProcessRequest](#gauge.messages.KillProcessRequest) |  | [KillProcessRequest](#gauge.messages.KillProcessRequest) |
| scenarioDataStoreInitRequest | [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest) |  | [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest) |
| specDataStoreInitRequest | [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest) |  | [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest) |
| suiteDataStoreInitRequest | [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest) |  | [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest) |
| stepNameRequest | [StepNameRequest](#gauge.messages.StepNameRequest) |  | [StepNameRequest](#gauge.messages.StepNameRequest) |
| stepNameResponse | [StepNameResponse](#gauge.messages.StepNameResponse) |  | [StepNameResponse](#gauge.messages.StepNameResponse) |
| refactorRequest | [RefactorRequest](#gauge.messages.RefactorRequest) |  | [RefactorRequest](#gauge.messages.RefactorRequest) |
| refactorResponse | [RefactorResponse](#gauge.messages.RefactorResponse) |  | [RefactorResponse](#gauge.messages.RefactorResponse) |
| unsupportedMessageResponse | [UnsupportedMessageResponse](#gauge.messages.UnsupportedMessageResponse) |  | [UnsupportedMessageResponse](#gauge.messages.UnsupportedMessageResponse) |
| cacheFileRequest | [CacheFileRequest](#gauge.messages.CacheFileRequest) |  | [CacheFileRequest](#gauge.messages.CacheFileRequest) |
| stepPositionsRequest | [StepPositionsRequest](#gauge.messages.StepPositionsRequest) |  | [StepPositionsRequest](#gauge.messages.StepPositionsRequest) |
| stepPositionsResponse | [StepPositionsResponse](#gauge.messages.StepPositionsResponse) |  | [StepPositionsResponse](#gauge.messages.StepPositionsResponse) |
| implementationFileListRequest | [ImplementationFileListRequest](#gauge.messages.ImplementationFileListRequest) |  | [ImplementationFileListRequest](#gauge.messages.ImplementationFileListRequest) |
| implementationFileListResponse | [ImplementationFileListResponse](#gauge.messages.ImplementationFileListResponse) |  | [ImplementationFileListResponse](#gauge.messages.ImplementationFileListResponse) |
| stubImplementationCodeRequest | [StubImplementationCodeRequest](#gauge.messages.StubImplementationCodeRequest) |  | [StubImplementationCodeRequest](#gauge.messages.StubImplementationCodeRequest) |
| fileDiff | [FileDiff](#gauge.messages.FileDiff) |  | [FileDiff](#gauge.messages.FileDiff) |
| implementationFileGlobPatternRequest | [ImplementationFileGlobPatternRequest](#gauge.messages.ImplementationFileGlobPatternRequest) |  | [ImplementationFileGlobPatternRequest](#gauge.messages.ImplementationFileGlobPatternRequest) |
| implementationFileGlobPatternResponse | [ImplementationFileGlobPatternResponse](#gauge.messages.ImplementationFileGlobPatternResponse) |  | [ImplementationFileGlobPatternResponse](#gauge.messages.ImplementationFileGlobPatternResponse) |
| suiteExecutionResultItem | [SuiteExecutionResultItem](#gauge.messages.SuiteExecutionResultItem) |  | [SuiteExecutionResult ](#gauge.messages.SuiteExecutionResult ) |
| keepAlive | [KeepAlive](#gauge.messages.KeepAlive) |  | [KeepAlive ](#gauge.messages.KeepAlive ) |






<a name="gauge.messages.ParameterPosition"></a>

### ParameterPosition
Holds the new and old positions of a parameter.
/ Used when refactoring a Step.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oldPosition | [int32](#int32) |  |  |
| newPosition | [int32](#int32) |  |  |






<a name="gauge.messages.RefactorRequest"></a>

### RefactorRequest
Tells the runner to refactor the specified Step.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| oldStepValue | [ProtoStepValue](#gauge.messages.ProtoStepValue) |  | Old value, used to lookup Step to refactor |
| newStepValue | [ProtoStepValue](#gauge.messages.ProtoStepValue) |  | New value, the to-be value of Step being refactored. |
| paramPositions | [ParameterPosition](#gauge.messages.ParameterPosition) | repeated | Holds parameter positions of all parameters. Contains old and new parameter positions. |
| saveChanges | [bool](#bool) |  | If set to true, the refactored files should be saved to the file system before returning the response. |






<a name="gauge.messages.RefactorResponse"></a>

### RefactorResponse
Response of a RefactorRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  | Flag indicating the success of Refactor operation. |
| error | [string](#string) |  | Error message, valid only if Refactor wasn&#39;t successful |
| filesChanged | [string](#string) | repeated | List of files that were affected because of the refactoring. |
| fileChanges | [FileChanges](#gauge.messages.FileChanges) | repeated | List of file changes to be made to successfully achieve refactoring. |






<a name="gauge.messages.ScenarioDataStoreInitRequest"></a>

### ScenarioDataStoreInitRequest
Request runner to initialize Scenario DataStore
/ Scenario Datastore is reset after every Scenario execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.ScenarioExecutionEndingRequest"></a>

### ScenarioExecutionEndingRequest
Sent at end of Scenario Execution. Tells the runner to execute `after_scenario` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current scenario execution info. |
| scenarioResult | [ProtoScenarioResult](#gauge.messages.ProtoScenarioResult) |  |  |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.ScenarioExecutionStartingRequest"></a>

### ScenarioExecutionStartingRequest
Sent at start of Scenario Execution. Tells the runner to execute `before_scenario` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current sceanrio execution info. |
| scenarioResult | [ProtoScenarioResult](#gauge.messages.ProtoScenarioResult) |  | Holds the scenarion result in scenarion execution starting. / Some fields will not be populated before execution. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.ScenarioInfo"></a>

### ScenarioInfo
Contains details of the Scenario execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the current Scenario being executed. |
| isFailed | [bool](#bool) |  | Flag to indicate if the current Scenario execution failed. |
| tags | [string](#string) | repeated | Tags relevant to the current Scenario execution. |






<a name="gauge.messages.SpecDataStoreInitRequest"></a>

### SpecDataStoreInitRequest
Request runner to initialize Spec DataStore
/ Spec Datastore is reset after every Spec execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.SpecDetails"></a>

### SpecDetails



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| details | [SpecDetails.SpecDetail](#gauge.messages.SpecDetails.SpecDetail) | repeated | Holds a collection of Spec details. |






<a name="gauge.messages.SpecDetails.SpecDetail"></a>

### SpecDetails.SpecDetail



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spec | [ProtoSpec](#gauge.messages.ProtoSpec) |  | Holds a collection of Specs that are defined in the project. |
| parseErrors | [Error](#gauge.messages.Error) | repeated | Holds a collection of parse errors present in the above spec. |






<a name="gauge.messages.SpecExecutionEndingRequest"></a>

### SpecExecutionEndingRequest
Sent at end of Spec Execution. Tells the runner to execute `after_spec` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current spec execution info. |
| specResult | [ProtoSpecResult](#gauge.messages.ProtoSpecResult) |  | Holds the specs result in spec execution ending. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.SpecExecutionStartingRequest"></a>

### SpecExecutionStartingRequest
Sent at start of Spec Execution. Tells the runner to execute `before_spec` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current spec execution info. |
| specResult | [ProtoSpecResult](#gauge.messages.ProtoSpecResult) |  | Holds the specs result in spec execution starting. / Some fields will not be populated before execution. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.SpecInfo"></a>

### SpecInfo
Contains details of the Spec execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | Name of the current Spec being executed. |
| fileName | [string](#string) |  | Full File path containing the current Spec being executed. |
| isFailed | [bool](#bool) |  | Flag to indicate if the current Spec execution failed. |
| tags | [string](#string) | repeated | Tags relevant to the current Spec execution. |






<a name="gauge.messages.StepExecutionEndingRequest"></a>

### StepExecutionEndingRequest
Sent at end of Step Execution. Tells the runner to execute `after_step` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current step execution info. |
| stepResult | [ProtoStepResult](#gauge.messages.ProtoStepResult) |  | Holds step result in step execution ending. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.StepExecutionStartingRequest"></a>

### StepExecutionStartingRequest
Sent at start of Step Execution. Tells the runner to execute `before_step` hook.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| currentExecutionInfo | [ExecutionInfo](#gauge.messages.ExecutionInfo) |  | Holds the current step execution info. |
| stepResult | [ProtoStepResult](#gauge.messages.ProtoStepResult) |  | Holds step result in step execution starting. / Some fields will not be populated before execution. |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.StepInfo"></a>

### StepInfo
Contains details of the Step execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| step | [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest) |  | The current request to execute Step |
| isFailed | [bool](#bool) |  | Flag to indicate if the current Step execution failed. |
| stackTrace | [string](#string) |  | The current stack trace in case of failure |
| errorMessage | [string](#string) |  | The error message in case of failure |






<a name="gauge.messages.StepNameRequest"></a>

### StepNameRequest
Request for details on a Single Step.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepValue | [string](#string) |  | Step text to lookup the Step. / This is the parsed step value, i.e. with placeholders for parameters. |






<a name="gauge.messages.StepNameResponse"></a>

### StepNameResponse
Response to StepNameRequest.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isStepPresent | [bool](#bool) |  | Flag indicating if there is a match for the given Step Text. |
| stepName | [string](#string) | repeated | The Step name of the given step. |
| hasAlias | [bool](#bool) |  | Flag indicating if the given Step is an alias. |
| fileName | [string](#string) |  | File name in which the step implementation exists |
| span | [Span](#gauge.messages.Span) |  | Range of step |
| isExternal | [bool](#bool) |  | Flag indicating if the given Step defined in some external library and the source code can not be accessed. |






<a name="gauge.messages.StepNamesRequest"></a>

### StepNamesRequest
Requests Gauge to give all Step Names.






<a name="gauge.messages.StepNamesResponse"></a>

### StepNamesResponse
Response to StepNamesRequest


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| steps | [string](#string) | repeated | Collection of strings corresponding to Step texts. |






<a name="gauge.messages.StepPositionsRequest"></a>

### StepPositionsRequest
Request for find step positions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| filePath | [string](#string) |  | Get step positions for file path |






<a name="gauge.messages.StepPositionsResponse"></a>

### StepPositionsResponse
Response for find step positions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepPositions | [StepPositionsResponse.StepPosition](#gauge.messages.StepPositionsResponse.StepPosition) | repeated | Step Position |
| error | [string](#string) |  | Error message |






<a name="gauge.messages.StepPositionsResponse.StepPosition"></a>

### StepPositionsResponse.StepPosition
Step position for each step implementation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepValue | [string](#string) |  | Step Value |
| span | [Span](#gauge.messages.Span) |  | Range of step |






<a name="gauge.messages.StepValidateRequest"></a>

### StepValidateRequest
Request sent ot the runner to check if given Step is valid.
/ The runner should check if there is an implementation defined for the given Step Text.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepText | [string](#string) |  | The text is used to lookup Step implementation |
| numberOfParameters | [int32](#int32) |  | The number of paramters in the Step |
| stepValue | [ProtoStepValue](#gauge.messages.ProtoStepValue) |  | This is use to generate step implementation template |






<a name="gauge.messages.StepValidateResponse"></a>

### StepValidateResponse
Response of StepValidateRequest.
/ The runner tells the caller if the Request was valid,
/ i.e. an implementation exists for given Step text.
/ Returns an error message if it is an error response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isValid | [bool](#bool) |  |  |
| errorMessage | [string](#string) |  |  |
| errorType | [StepValidateResponse.ErrorType](#gauge.messages.StepValidateResponse.ErrorType) |  |  |
| suggestion | [string](#string) |  |  |






<a name="gauge.messages.StubImplementationCodeRequest"></a>

### StubImplementationCodeRequest
Request for injecting code snippet into implementation file


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| implementationFilePath | [string](#string) |  | Path of the file where the new stub implementation will be added |
| codes | [string](#string) | repeated | List of implementation codes to be appended to implementation file. |






<a name="gauge.messages.SuiteDataStoreInitRequest"></a>

### SuiteDataStoreInitRequest
Request runner to initialize Suite DataStore
/ Suite Datastore is reset after every Suite execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stream | [int32](#int32) |  |  |






<a name="gauge.messages.SuiteExecutionResult"></a>

### SuiteExecutionResult
Result of the Suite Execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| suiteResult | [ProtoSuiteResult](#gauge.messages.ProtoSuiteResult) |  |  |






<a name="gauge.messages.SuiteExecutionResultItem"></a>

### SuiteExecutionResultItem



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| resultItem | [ProtoItem](#gauge.messages.ProtoItem) |  |  |






<a name="gauge.messages.TextDiff"></a>

### TextDiff
A Single Replace Diff Element to be applied


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| span | [Span](#gauge.messages.Span) |  | Range of file to be replaced |
| content | [string](#string) |  | New content to replace the content in the span |






<a name="gauge.messages.UnsupportedMessageResponse"></a>

### UnsupportedMessageResponse
Response when a unsupported message request is sent.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |





 


<a name="gauge.messages.CacheFileRequest.FileStatus"></a>

### CacheFileRequest.FileStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| CHANGED | 0 | The file content was changed in the client |
| CLOSED | 1 | The file was closed in the client |
| CREATED | 2 | The file was created on the client |
| DELETED | 3 | The file was deleted on the client |
| OPENED | 4 | The file is opened in the client |



<a name="gauge.messages.Message.MessageType"></a>

### Message.MessageType


| Name | Number | Description |
| ---- | ------ | ----------- |
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
| UnsupportedMessageResponse | 23 |  |
| CacheFileRequest | 24 |  |
| StepPositionsRequest | 25 |  |
| StepPositionsResponse | 26 |  |
| ImplementationFileListRequest | 27 |  |
| ImplementationFileListResponse | 28 |  |
| StubImplementationCodeRequest | 29 |  |
| FileDiff | 30 |  |
| ImplementationFileGlobPatternRequest | 31 |  |
| ImplementationFileGlobPatternResponse | 32 |  |
| SuiteExecutionResultItem | 33 |  |
| KeepAlive | 34 |  |



<a name="gauge.messages.StepValidateResponse.ErrorType"></a>

### StepValidateResponse.ErrorType


| Name | Number | Description |
| ---- | ------ | ----------- |
| STEP_IMPLEMENTATION_NOT_FOUND | 0 |  |
| DUPLICATE_STEP_IMPLEMENTATION | 1 |  |


 

 

 



<a name="services.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## services.proto


 

 

 


<a name="gauge.messages.Documenter"></a>

### Documenter
Reporter services is meant for documentation plugins

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GenerateDocs | [SpecDetails](#gauge.messages.SpecDetails) | [Empty](#gauge.messages.Empty) | GenerateDocs is a RPC tell plugin to generate docs from the spec details.

Accepts a SpecDetails message and returns a Empty message. |
| Kill | [KillProcessRequest](#gauge.messages.KillProcessRequest) | [Empty](#gauge.messages.Empty) | Kill is a RPC tell plugin to stop grpc server and kill the plugin process.

Accepts a KillProcessRequest message and returns a Empty message. |


<a name="gauge.messages.Reporter"></a>

### Reporter
Reporter services is meant for reporting plugins, or others plugins which are interested the live events

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| NotifyExecutionStarting | [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest) | [Empty](#gauge.messages.Empty) | NotifyExecutionStarting is a RPC to tell plugins that the execution has started. Accepts a ExecutionStartingRequest message and returns a Empty message |
| NotifySpecExecutionStarting | [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest) | [Empty](#gauge.messages.Empty) | NotifySpecExecutionStarting is a RPC to tell plugins that the specification execution has started. Accepts a SpecExecutionStartingRequest message and returns a Empty message |
| NotifyScenarioExecutionStarting | [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest) | [Empty](#gauge.messages.Empty) | NotifyScenarioExecutionStarting is a RPC to tell plugins that the scenario execution has started. Accepts a ScenarioExecutionStartingRequest message and returns a Empty message |
| NotifyStepExecutionStarting | [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest) | [Empty](#gauge.messages.Empty) | NotifyStepExecutionStarting is a RPC to tell plugins that the step execution has started. Accepts a StepExecutionStartingRequest message and returns a Empty message |
| NotifyStepExecutionEnding | [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest) | [Empty](#gauge.messages.Empty) | NotifyStepExecutionEnding is a RPC to tell plugins that the step execution has finished. Accepts a StepExecutionEndingRequest message and returns a Empty message |
| NotifyScenarioExecutionEnding | [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest) | [Empty](#gauge.messages.Empty) | NotifyScenarioExecutionEnding is a RPC to tell plugins that the scenario execution has finished. Accepts a ScenarioExecutionEndingRequest message and returns a Empty message |
| NotifySpecExecutionEnding | [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest) | [Empty](#gauge.messages.Empty) | NotifySpecExecutionEnding is a RPC to tell plugins that the specification execution has finished. Accepts a SpecExecutionEndingRequest message and returns a Empty message |
| NotifyExecutionEnding | [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest) | [Empty](#gauge.messages.Empty) | NotifyExecutionEnding is a RPC to tell plugins that the execution has finished. Accepts a ExecutionEndingRequest message and returns a Empty message |
| NotifySuiteResult | [SuiteExecutionResult](#gauge.messages.SuiteExecutionResult) | [Empty](#gauge.messages.Empty) | NotifySuiteResult is a RPC to tell about the end result of execution Accepts a SuiteExecutionResult message and returns a Empty message. |
| Kill | [KillProcessRequest](#gauge.messages.KillProcessRequest) | [Empty](#gauge.messages.Empty) | Kill is a RPC tell plugin to stop grpc server and kill the plugin process. Accepts a KillProcessRequest message and returns a Empty message. |


<a name="gauge.messages.Runner"></a>

### Runner


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ValidateStep | [StepValidateRequest](#gauge.messages.StepValidateRequest) | [StepValidateResponse](#gauge.messages.StepValidateResponse) | ValidateStep is a RPC to validate a given step. Accepts a StepValidateRequest message and returns a StepValidateResponse message |
| InitializeSuiteDataStore | [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | SuiteDataStoreInit is a RPC to initialize the suite level data store. Accepts a Empty message and returns a ExecutionStatusResponse message |
| StartExecution | [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | ExecutionStarting is a RPC to tell runner to execute Suite level hooks. Accepts a ExecutionStartingRequest message and returns a ExecutionStatusResponse message |
| InitializeSpecDataStore | [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | SpecDataStoreInit is a RPC to initialize the spec level data store. Accepts a Empty message and returns a ExecutionStatusResponse message |
| StartSpecExecution | [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | SpecExecutionStarting is a RPC to tell runner to execute spec level hooks. Accepts a SpecExecutionStartingRequest message and returns a ExecutionStatusResponse message |
| InitializeScenarioDataStore | [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | ScenarioDataStoreInit is a RPC to initialize the scenario level data store. Accepts a Empty message and returns a ExecutionStatusResponse message |
| StartScenarioExecution | [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | ScenarioExecutionStarting is a RPC to tell runner to execute scenario level hooks. Accepts a ScenarioExecutionStartingRequest message and returns a ExecutionStatusResponse message |
| StartStepExecution | [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | StepExecutionStarting is a RPC to tell runner to execute step level hooks. Accepts a StepExecutionStartingRequest message and returns a ExecutionStatusResponse message |
| ExecuteStep | [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | ExecuteStep is a RPC to tell runner to execute a step. Accepts a ExecuteStepRequest message and returns a ExecutionStatusResponse message |
| FinishStepExecution | [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | StepExecutionEnding is a RPC to tell runner to execute step level hooks. Accepts a StepExecutionEndingRequest message and returns a ExecutionStatusResponse message |
| FinishScenarioExecution | [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | ScenarioExecutionEnding is a RPC to tell runner to execute Scenario level hooks. Accepts a ScenarioExecutionEndingRequest message and returns a ExecutionStatusResponse message |
| FinishSpecExecution | [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | SpecExecutionEnding is a RPC to tell runner to execute spec level hooks. Accepts a SpecExecutionEndingRequest message and returns a ExecutionStatusResponse message |
| FinishExecution | [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest) | [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse) | ExecutionEnding is a RPC to tell runner to execute suite level hooks. Accepts a ExecutionEndingRequest message and returns a ExecutionStatusResponse message |
| CacheFile | [CacheFileRequest](#gauge.messages.CacheFileRequest) | [Empty](#gauge.messages.Empty) | CacheFile is a RPC to tell runner to load/reload/unload a implementation file. Accepts a CacheFileRequest message and returns a Empty message |
| GetStepName | [StepNameRequest](#gauge.messages.StepNameRequest) | [StepNameResponse](#gauge.messages.StepNameResponse) | GetStepName is a RPC to get information about the given step. Accepts a StepNameRequest message and returns a StepNameResponse message. |
| GetGlobPatterns | [Empty](#gauge.messages.Empty) | [ImplementationFileGlobPatternResponse](#gauge.messages.ImplementationFileGlobPatternResponse) | GetGlobPatterns is a RPC to get the file path pattern which needs to be cached. Accepts a Empty message and returns a ImplementationFileGlobPatternResponse message. |
| GetStepNames | [StepNamesRequest](#gauge.messages.StepNamesRequest) | [StepNamesResponse](#gauge.messages.StepNamesResponse) | GetStepNames is a RPC to get all the available steps from the runner. Accepts a StepNamesRequest message and returns a StepNamesResponse |
| GetStepPositions | [StepPositionsRequest](#gauge.messages.StepPositionsRequest) | [StepPositionsResponse](#gauge.messages.StepPositionsResponse) | GetStepPositions is a RPC to get positions of all available steps in a given file. Accepts a StepPositionsRequest message and returns a StepPositionsResponse message |
| GetImplementationFiles | [Empty](#gauge.messages.Empty) | [ImplementationFileListResponse](#gauge.messages.ImplementationFileListResponse) | GetImplementationFiles is a RPC get all the existing implementation files. Accepts a Empty and returns a ImplementationFileListResponse message. |
| ImplementStub | [StubImplementationCodeRequest](#gauge.messages.StubImplementationCodeRequest) | [FileDiff](#gauge.messages.FileDiff) | ImplementStub is a RPC to to ask runner to add a given implementation to given file. Accepts a StubImplementationCodeRequest and returns a FileDiff message. |
| Refactor | [RefactorRequest](#gauge.messages.RefactorRequest) | [RefactorResponse](#gauge.messages.RefactorResponse) | Refactor is a RPC to refactor a given step in implementation file. Accepts a RefactorRequest message and returns a RefactorResponse message. |
| Kill | [KillProcessRequest](#gauge.messages.KillProcessRequest) | [Empty](#gauge.messages.Empty) | Kill is a RPC tell plugin to stop grpc server and kill the plugin process. Accepts a KillProcessRequest message and returns a Empty message. |

 



<a name="spec.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## spec.proto



<a name="gauge.messages.Error"></a>

### Error
A proto object representing an error in spec/Scenario.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [Error.ErrorType](#gauge.messages.Error.ErrorType) |  | Holds the type of error |
| filename | [string](#string) |  | Holds the filename. |
| lineNumber | [int32](#int32) |  | Holds the line number of the error in file. |
| message | [string](#string) |  | Holds the error message. |






<a name="gauge.messages.Fragment"></a>

### Fragment
A proto object representing Fragment.
/ Fragments, put together make up A Step


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fragmentType | [Fragment.FragmentType](#gauge.messages.Fragment.FragmentType) |  | Type of Fragment, valid values are Text, Parameter |
| text | [string](#string) |  | Text part of the Fragment, valid only if FragmentType=Text |
| parameter | [Parameter](#gauge.messages.Parameter) |  | Parameter part of the Fragment, valid only if FragmentType=Parameter |






<a name="gauge.messages.Parameter"></a>

### Parameter
A proto object representing Fragment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parameterType | [Parameter.ParameterType](#gauge.messages.Parameter.ParameterType) |  | Type of the Parameter. Valid values: Static, Dynamic, Special_String, Special_Table, Table |
| value | [string](#string) |  | Holds the value of the parameter |
| name | [string](#string) |  | Holds the name of the parameter, used as Key to lookup the value. |
| table | [ProtoTable](#gauge.messages.ProtoTable) |  | Holds the table value, if parameterType=Table or Special_Table |






<a name="gauge.messages.ProtoComment"></a>

### ProtoComment
A proto object representing Comment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  | Text representing the Comment. |






<a name="gauge.messages.ProtoConcept"></a>

### ProtoConcept
Concept is a type of step, that can have multiple Steps.
/ But from a caller&#39;s perspective, it is still used as any other Step
/ A proto object representing a Concept


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| conceptStep | [ProtoStep](#gauge.messages.ProtoStep) |  | Represents the Step value of a Concept. |
| steps | [ProtoItem](#gauge.messages.ProtoItem) | repeated | Collection of Steps in the given concepts. |
| conceptExecutionResult | [ProtoStepExecutionResult](#gauge.messages.ProtoStepExecutionResult) |  | Holds the execution result. |






<a name="gauge.messages.ProtoExecutionResult"></a>

### ProtoExecutionResult
A proto object representing the result of an execution


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| failed | [bool](#bool) |  | Flag to indicate failure |
| recoverableError | [bool](#bool) |  | Flag to indicate if the error is recoverable from. |
| errorMessage | [string](#string) |  | The actual error message. |
| stackTrace | [string](#string) |  | Stacktrace of the error |
| screenShot | [bytes](#bytes) |  | [DEPRECATED, use failureScreenshotFile] Bytes containing screenshot taken at the time of failure. |
| executionTime | [int64](#int64) |  | Holds the time taken for executing this scenario. |
| message | [string](#string) | repeated | Additional information at exec time to be available on reports |
| errorType | [ProtoExecutionResult.ErrorType](#gauge.messages.ProtoExecutionResult.ErrorType) |  | Type of the Error. Valid values: ASSERTION, VERIFICATION. Default: ASSERTION |
| failureScreenshot | [bytes](#bytes) |  | [DEPRECATED, use failureScreenshotFile] Bytes containing screenshot taken at the time of failure. |
| screenshots | [bytes](#bytes) | repeated | [DEPRECATED, use screenshotFiles] Bytes array containing screenshots at the time of it invoked |
| failureScreenshotFile | [string](#string) |  | Path to the screenshot file captured at the time of failure. |
| screenshotFiles | [string](#string) | repeated | Path to the screenshot files captured using Gauge screenshsot API. |






<a name="gauge.messages.ProtoHookFailure"></a>

### ProtoHookFailure
A proto object representing a pre-hook failure.
/ Used to hold failure information for before_suite, before_spec, before_scenario and before_spec hooks.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stackTrace | [string](#string) |  | Stacktrace from the failure |
| errorMessage | [string](#string) |  | Error message from the failure |
| screenShot | [bytes](#bytes) |  | [DEPRECATED, use failureScreenshotFile] Bytes holding the screenshot taken at the time of failure. |
| tableRowIndex | [int32](#int32) |  | Contains table row index corresponding to datatable rows |
| failureScreenshot | [bytes](#bytes) |  | [DEPRECATED, use failureScreenshotFile] Bytes holding the screenshot taken at the time of failure. |
| failureScreenshotFile | [string](#string) |  | Path to the screenshot file captured at the time of failure. |






<a name="gauge.messages.ProtoItem"></a>

### ProtoItem
Container for all valid Items under a Specification.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| itemType | [ProtoItem.ItemType](#gauge.messages.ProtoItem.ItemType) |  | Itemtype of the current ProtoItem |
| step | [ProtoStep](#gauge.messages.ProtoStep) |  | Holds the Step definition. Valid only if ItemType = Step |
| concept | [ProtoConcept](#gauge.messages.ProtoConcept) |  | Holds the Concept definition. Valid only if ItemType = Concept |
| scenario | [ProtoScenario](#gauge.messages.ProtoScenario) |  | Holds the Scenario definition. Valid only if ItemType = Scenario |
| tableDrivenScenario | [ProtoTableDrivenScenario](#gauge.messages.ProtoTableDrivenScenario) |  | Holds the TableDrivenScenario definition. Valid only if ItemType = TableDrivenScenario |
| comment | [ProtoComment](#gauge.messages.ProtoComment) |  | Holds the Comment definition. Valid only if ItemType = Comment |
| table | [ProtoTable](#gauge.messages.ProtoTable) |  | Holds the Table definition. Valid only if ItemType = Table |
| tags | [ProtoTags](#gauge.messages.ProtoTags) |  | Holds the Tags definition. Valid only if ItemType = Tags |
| fileName | [string](#string) |  | Holds the Filename that the item belongs to |






<a name="gauge.messages.ProtoScenario"></a>

### ProtoScenario
A proto object representing a Scenario


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scenarioHeading | [string](#string) |  | Heading of the given Scenario |
| failed | [bool](#bool) |  | Flag to indicate if the Scenario execution failed |
| contexts | [ProtoItem](#gauge.messages.ProtoItem) | repeated | Collection of Context steps. The Context steps are executed before every run. |
| scenarioItems | [ProtoItem](#gauge.messages.ProtoItem) | repeated | Collection of Items under a scenario. These could be Steps, Comments, Tags, TableDrivenScenarios or Tables |
| preHookFailure | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) |  | Contains a &#39;before&#39; hook failure message. This happens when the `before_scenario` hook has an error. |
| postHookFailure | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) |  | Contains a &#39;after&#39; hook failure message. This happens when the `after_scenario` hook has an error. |
| tags | [string](#string) | repeated | Contains a list of tags that are defined at the specification level. Scenario tags are not present here. |
| executionTime | [int64](#int64) |  | Holds the time taken for executing this scenario. |
| skipped | [bool](#bool) |  | Flag to indicate if the Scenario execution is skipped |
| skipErrors | [string](#string) | repeated | Holds the error messages for skipping scenario from execution |
| ID | [string](#string) |  | Holds the unique Identifier of a scenario. |
| tearDownSteps | [ProtoItem](#gauge.messages.ProtoItem) | repeated | Collection of Teardown steps. The Teardown steps are executed after every run. |
| span | [Span](#gauge.messages.Span) |  | Span(start, end) of scenario |
| executionStatus | [ExecutionStatus](#gauge.messages.ExecutionStatus) |  | Execution status for the scenario |
| preHookMessages | [string](#string) | repeated | Additional information at pre hook exec time to be available on reports |
| postHookMessages | [string](#string) | repeated | Additional information at post hook exec time to be available on reports |
| preHookMessage | [string](#string) | repeated | [DEPRECATED, use preHookMessages] Additional information at pre hook exec time to be available on reports |
| postHookMessage | [string](#string) | repeated | [DEPRECATED, use postHookMessages] Additional information at post hook exec time to be available on reports |
| preHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use preHookScreenshotFiles] Capture Screenshot at pre hook exec time to be available on reports |
| postHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use postHookScreenshotFiles] Capture Screenshot at post hook exec time to be available on reports |
| preHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on pre hook exec time to be available on reports |
| postHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on post hook exec time to be available on reports |
| retriesCount | [int64](#int64) |  | Number of retires |






<a name="gauge.messages.ProtoScenarioResult"></a>

### ProtoScenarioResult
A proto object representing the result of Scenario execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protoItem | [ProtoItem](#gauge.messages.ProtoItem) |  | Collection of scenarios in scenario execution result. |
| executionTime | [int64](#int64) |  | Holds the time taken for executing the whole suite. |
| timestamp | [string](#string) |  | Holds the timestamp of event starting. |






<a name="gauge.messages.ProtoSpec"></a>

### ProtoSpec
A proto object representing a Specification
/ A specification can contain Scenarios or Steps, besides Comments


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specHeading | [string](#string) |  | Heading describing the Specification |
| items | [ProtoItem](#gauge.messages.ProtoItem) | repeated | A collection of items that come under this step |
| isTableDriven | [bool](#bool) |  | Flag indicating if this is a Table Driven Specification. The table is defined in the context, this is different from using a table parameter. |
| preHookFailures | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) | repeated | Contains a &#39;before&#39; hook failure message. This happens when the `before_spec` hook has an error. |
| postHookFailures | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) | repeated | Contains a &#39;before&#39; hook failure message. This happens when the `after_hook` hook has an error. |
| fileName | [string](#string) |  | Contains the filename for that holds this specification. |
| tags | [string](#string) | repeated | Contains a list of tags that are defined at the specification level. Scenario tags are not present here. |
| preHookMessages | [string](#string) | repeated | Additional information at pre hook exec time to be available on reports |
| postHookMessages | [string](#string) | repeated | Additional information at post hook exec time to be available on reports |
| preHookMessage | [string](#string) | repeated | [DEPRECATED, use preHookMessages] Additional information at pre hook exec time to be available on reports |
| postHookMessage | [string](#string) | repeated | [DEPRECATED, use postHookMessages] Additional information at post hook exec time to be available on reports |
| preHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use preHookScreenshotFiles] Capture Screenshot at pre hook exec time to be available on reports |
| postHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use postHookScreenshotFiles] Capture Screenshot at post hook exec time to be available on reports |
| itemCount | [int64](#int64) |  | meta field to indicate the number of items in the list / used when items are sent as individual chunk |
| preHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on pre hook exec time to be available on reports |
| postHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on post hook exec time to be available on reports |






<a name="gauge.messages.ProtoSpecResult"></a>

### ProtoSpecResult
A proto object representing the result of Spec execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protoSpec | [ProtoSpec](#gauge.messages.ProtoSpec) |  | Represents the corresponding Specification |
| scenarioCount | [int32](#int32) |  | Holds the number of Scenarios executed |
| scenarioFailedCount | [int32](#int32) |  | Holds the number of Scenarios failed |
| failed | [bool](#bool) |  | Flag to indicate failure |
| failedDataTableRows | [int32](#int32) | repeated | Holds the row numbers, which caused the execution to fail. |
| executionTime | [int64](#int64) |  | Holds the time taken for executing the spec. |
| skipped | [bool](#bool) |  | Flag to indicate if spec is skipped |
| scenarioSkippedCount | [int32](#int32) |  | Holds the number of Scenarios skipped |
| skippedDataTableRows | [int32](#int32) | repeated | Holds the row numbers, for which the execution skipped. |
| errors | [Error](#gauge.messages.Error) | repeated | Holds parse, validation and skipped errors. |
| timestamp | [string](#string) |  | Holds the timestamp of event starting. |






<a name="gauge.messages.ProtoStep"></a>

### ProtoStep
A proto object representing a Step


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actualText | [string](#string) |  | Holds the raw text of the Step as defined in the spec file. This contains the actual parameter values. |
| parsedText | [string](#string) |  | Contains the parsed text of the Step. This will have placeholders for the parameters. |
| fragments | [Fragment](#gauge.messages.Fragment) | repeated | Collection of a list of fragments for a Step. A fragment could be either text or parameter. |
| stepExecutionResult | [ProtoStepExecutionResult](#gauge.messages.ProtoStepExecutionResult) |  | Holds the result from the execution. |
| preHookMessages | [string](#string) | repeated | Additional information at pre hook exec time to be available on reports |
| postHookMessages | [string](#string) | repeated | Additional information at post hook exec time to be available on reports |
| preHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use preHookScreenshotFiles] Capture Screenshot at pre hook exec time to be available on reports |
| postHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use postHookScreenshotFiles] Capture Screenshot at post hook exec time to be available on reports |
| preHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on pre hook exec time to be available on reports |
| postHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on post hook exec time to be available on reports |






<a name="gauge.messages.ProtoStepExecutionResult"></a>

### ProtoStepExecutionResult
A proto object representing Step Execution result


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| executionResult | [ProtoExecutionResult](#gauge.messages.ProtoExecutionResult) |  | The actual result of the execution |
| preHookFailure | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) |  | Contains a &#39;before&#39; hook failure message. This happens when the `before_step` hook has an error. |
| postHookFailure | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) |  | Contains a &#39;after&#39; hook failure message. This happens when the `after_step` hook has an error. |
| skipped | [bool](#bool) |  |  |
| skippedReason | [string](#string) |  |  |






<a name="gauge.messages.ProtoStepResult"></a>

### ProtoStepResult
A proto object representing the result of Step execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| protoItem | [ProtoItem](#gauge.messages.ProtoItem) |  | Collection of steps in step execution result. |
| executionTime | [int64](#int64) |  | Holds the time taken for executing the whole suite. |
| timestamp | [string](#string) |  | Holds the timestamp of event starting. |






<a name="gauge.messages.ProtoStepValue"></a>

### ProtoStepValue
A proto object representing a Step value.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| stepValue | [string](#string) |  | The actual string value describing he Step |
| parameterizedStepValue | [string](#string) |  | The parameterized string value describing he Step. The parameters are replaced with placeholders. |
| parameters | [string](#string) | repeated | A collection of strings representing the parameters. |






<a name="gauge.messages.ProtoSuiteResult"></a>

### ProtoSuiteResult
A proto object representing the result of entire Suite execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| specResults | [ProtoSpecResult](#gauge.messages.ProtoSpecResult) | repeated | Contains the result from the execution |
| preHookFailure | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) |  | Contains a &#39;before&#39; hook failure message. This happens when the `before_suite` hook has an error |
| postHookFailure | [ProtoHookFailure](#gauge.messages.ProtoHookFailure) |  | Contains a &#39;after&#39; hook failure message. This happens when the `after_suite` hook has an error |
| failed | [bool](#bool) |  | Flag to indicate failure |
| specsFailedCount | [int32](#int32) |  | Holds the count of number of Specifications that failed. |
| executionTime | [int64](#int64) |  | Holds the time taken for executing the whole suite. |
| successRate | [float](#float) |  | Holds a metric indicating the success rate of the execution. |
| environment | [string](#string) |  | The environment against which execution was done |
| tags | [string](#string) |  | Tag expression used for filtering specification |
| projectName | [string](#string) |  | Project name |
| timestamp | [string](#string) |  | Timestamp of when execution started |
| specsSkippedCount | [int32](#int32) |  |  |
| preHookMessages | [string](#string) | repeated | Additional information at pre hook exec time to be available on reports |
| postHookMessages | [string](#string) | repeated | Additional information at post hook exec time to be available on reports |
| preHookMessage | [string](#string) | repeated | [DEPRECATED, use preHookMessages] Additional information at pre hook exec time to be available on reports |
| postHookMessage | [string](#string) | repeated | [DEPRECATED, use postHookMessages] Additional information at post hook exec time to be available on reports |
| preHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use preHookScreenshotFiles] Capture Screenshot at pre hook exec time to be available on reports |
| postHookScreenshots | [bytes](#bytes) | repeated | [DEPRECATED, use postHookScreenshotFiles] Capture Screenshot at post hook exec time to be available on reports |
| chunked | [bool](#bool) |  | Indicates if the result is sent in chunks |
| chunkSize | [int64](#int64) |  | Indicates the number of chunks to expect after this |
| preHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on pre hook exec time to be available on reports |
| postHookScreenshotFiles | [string](#string) | repeated | Screenshots captured on post hook exec time to be available on reports |






<a name="gauge.messages.ProtoTable"></a>

### ProtoTable
A proto object representing Table.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| headers | [ProtoTableRow](#gauge.messages.ProtoTableRow) |  | Contains the Headers for the table |
| rows | [ProtoTableRow](#gauge.messages.ProtoTableRow) | repeated | Contains the Rows for the table |






<a name="gauge.messages.ProtoTableDrivenScenario"></a>

### ProtoTableDrivenScenario
A proto object representing a TableDrivenScenario


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| scenario | [ProtoScenario](#gauge.messages.ProtoScenario) |  | Scenario under Table driven execution |
| tableRowIndex | [int32](#int32) |  | Row Index of data table against which the current scenario is executed |
| scenarioTableRowIndex | [int32](#int32) |  | Row Index of scenario data table against which the current scenario is executed |
| isSpecTableDriven | [bool](#bool) |  | Executed against a spec data table |
| isScenarioTableDriven | [bool](#bool) |  | Executed against a scenario data table |
| scenarioDataTable | [ProtoTable](#gauge.messages.ProtoTable) |  | Holds the scenario data table |
| scenarioTableRow | [ProtoTable](#gauge.messages.ProtoTable) |  | Hold the row of scenario data table. |






<a name="gauge.messages.ProtoTableRow"></a>

### ProtoTableRow
A proto object representing Table.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cells | [string](#string) | repeated | Represents the cells of a given table |






<a name="gauge.messages.ProtoTags"></a>

### ProtoTags
A proto object representing Tags


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tags | [string](#string) | repeated | A collection of Tags |






<a name="gauge.messages.Span"></a>

### Span
A proto object representing a Span of content


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start | [int64](#int64) |  |  |
| end | [int64](#int64) |  |  |
| startChar | [int64](#int64) |  |  |
| endChar | [int64](#int64) |  |  |





 


<a name="gauge.messages.Error.ErrorType"></a>

### Error.ErrorType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PARSE_ERROR | 0 |  |
| VALIDATION_ERROR | 1 |  |



<a name="gauge.messages.ExecutionStatus"></a>

### ExecutionStatus
Execution Status

| Name | Number | Description |
| ---- | ------ | ----------- |
| NOTEXECUTED | 0 |  |
| PASSED | 1 |  |
| FAILED | 2 |  |
| SKIPPED | 3 |  |



<a name="gauge.messages.Fragment.FragmentType"></a>

### Fragment.FragmentType
Enum representing the types of Fragment

| Name | Number | Description |
| ---- | ------ | ----------- |
| Text | 0 | Fragment is a Text part |
| Parameter | 1 | Fragment is a Parameter part |



<a name="gauge.messages.Parameter.ParameterType"></a>

### Parameter.ParameterType
Enum representing types of Parameter.

| Name | Number | Description |
| ---- | ------ | ----------- |
| Static | 0 | Static parameter. The value of the parameter is defined at the Step. |
| Dynamic | 1 | Dynamic parameter. This is a parameter placeholder, and the actual value is injected at runtime, depending on the context of the call. |
| Special_String | 2 | Special paramter, taking a string value. Special paramters are read from a file. |
| Special_Table | 3 | Special paramter, taking a Table value. This parameter is read from a csv file. |
| Table | 4 | A table parameter, used for data driven execution. |



<a name="gauge.messages.ProtoExecutionResult.ErrorType"></a>

### ProtoExecutionResult.ErrorType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ASSERTION | 0 |  |
| VERIFICATION | 1 |  |



<a name="gauge.messages.ProtoItem.ItemType"></a>

### ProtoItem.ItemType
Enumerates various item types that the proto item can contain. Valid types are: Step, Comment, Concept, Scenario, TableDrivenScenario, Table, Tags

| Name | Number | Description |
| ---- | ------ | ----------- |
| Step | 0 | Item is a Step |
| Comment | 1 | Item is a Comment |
| Concept | 2 | Item is a Concept |
| Scenario | 3 | Item is a Scenario |
| TableDrivenScenario | 4 | Item is a TableDrivenScenario, a special case of Scenario, where there is a Context Step defining a table. |
| Table | 5 | Item is a Table |
| Tags | 6 | Item is a Tag |


 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

