package type_event

type EventChangeTag struct {
	Event      `xml:",omitempty,inline"`
	ChangeType string `xml:"ChangeType"`
}

/*
<xml>
	<ToUserName><![CDATA[CORPID]]></ToUserName>
	<FromUserName><![CDATA[sys]]></FromUserName>
	<CreateTime>1403610513</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[change_external_tag]]></Event>
	<Id><![CDATA[TAG_ID]]></Id>
	<TagType><![CDATA[tag]]></TagType>
	<ChangeType><![CDATA[create]]></ChangeType>
	<StrategyId>1</StrategyId>
</xml>
*/

// ChangeExternalTagCreate 企业客户标签创建事件
type ChangeExternalTagCreate struct {
	EventChangeTag `xml:",omitempty,inline"`
	Id             string `xml:"Id"`
	TagType        string `xml:"TagType"`
	StrategyId     string `xml:"StrategyId"`
}

/*
<xml>
	<ToUserName><![CDATA[CORPID]]></ToUserName>
	<FromUserName><![CDATA[sys]]></FromUserName>
	<CreateTime>1403610513</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[change_external_tag]]></Event>
	<Id><![CDATA[TAG_ID]]></Id>
	<TagType><![CDATA[tag]]></TagType>
	<ChangeType><![CDATA[update]]></ChangeType>
	<StrategyId>1</StrategyId>
</xml>
*/

// ChangeExternalTagCreate 企业客户标签更新事件
type ChangeExternalTagUpdate struct {
	EventChangeTag `xml:",omitempty,inline"`
	Id             string `xml:"Id"`
	TagType        string `xml:"TagType"`
	StrategyId     string `xml:"StrategyId"`
}

/*
<xml>
	<ToUserName><![CDATA[CORPID]]></ToUserName>
	<FromUserName><![CDATA[sys]]></FromUserName>
	<CreateTime>1403610513</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[change_external_tag]]></Event>
	<Id><![CDATA[TAG_ID]]></Id>
	<TagType><![CDATA[tag]]></TagType>
	<ChangeType><![CDATA[delete]]></ChangeType>
	<StrategyId>1</StrategyId>
</xml>
*/

// ChangeExternalTagCreate 企业客户标签删除事件
type ChangeExternalTagDelete struct {
	EventChangeTag `xml:",omitempty,inline"`
	Id             string `xml:"Id"`
	TagType        string `xml:"TagType"`
	StrategyId     string `xml:"StrategyId"`
}

/*
<xml>
	<ToUserName><![CDATA[CORPID]]></ToUserName>
	<FromUserName><![CDATA[sys]]></FromUserName>
	<CreateTime>1403610513</CreateTime>
	<MsgType><![CDATA[event]]></MsgType>
	<Event><![CDATA[change_external_tag]]></Event>
	<Id><![CDATA[TAG_ID]]></Id>
	<StrategyId><![CDATA[STRATEGY_ID]]></StrategyId>
	<ChangeType><![CDATA[shuffle]]></ChangeType>
</xml>
*/

// ChangeExternalTagCreate 企业客户标签重排事件
type ChangeExternalTagShuffle struct {
	EventChangeTag `xml:",omitempty,inline"`
	Id             string `xml:"Id"`
	TagType        string `xml:"TagType"`
	StrategyId     string `xml:"StrategyId"`
}
