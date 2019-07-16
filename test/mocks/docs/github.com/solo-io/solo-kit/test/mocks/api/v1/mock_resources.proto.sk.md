<!-- Code generated by solo-kit. DO NOT EDIT. -->

### Package: `testing.solo.io`  
Syntax Comments
Syntax Comments a


 
#### Types:


- [MockResource](#MockResource) **Top-Level Resource**
- [FakeResource](#FakeResource) **Top-Level Resource**
- [MockXdsResourceConfig](#MockXdsResourceConfig)
  



##### Source File: [github.com/solo-io/solo-kit/test/mocks/api/v1/mock_resources.proto](https://github.com/solo-io/solo-kit/blob/master/test/mocks/api/v1/mock_resources.proto)





---
### <a name=MockResource>MockResource</a>

 
Mock resources for goofin off

```yaml
"status": .core.solo.io.Status
"metadata": .core.solo.io.Metadata
"data": string
"someDumbField": string
"oneofOne": string
"oneofTwo": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `status` | [.core.solo.io.Status](../../../../api/v1/status.proto.sk.md#Status) |  |  |
| `metadata` | [.core.solo.io.Metadata](../../../../api/v1/metadata.proto.sk.md#Metadata) |  |  |
| `data` | `string` |  |  |
| `someDumbField` | `string` |  |  |
| `oneofOne` | `string` |  |  |
| `oneofTwo` | `bool` |  |  |




---
### <a name=FakeResource>FakeResource</a>



```yaml
"count": int
"metadata": .core.solo.io.Metadata
"status": .core.solo.io.Status

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `count` | `int` |  |  |
| `metadata` | [.core.solo.io.Metadata](../../../../api/v1/metadata.proto.sk.md#Metadata) |  |  |
| `status` | [.core.solo.io.Status](../../../../api/v1/status.proto.sk.md#Status) |  |  |




---
### <a name=MockXdsResourceConfig>MockXdsResourceConfig</a>

 


```yaml
"domain": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `domain` | `string` |  |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->