// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10
package PaymentReturn_pacs_004_001_10

import (
	"encoding/xml"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	PmtRtr PaymentReturnV10 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PmtRtr"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	IBAN *IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Othr,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    *ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type ActiveCurrencyAndAmountFedwire1 struct {
	Value ActiveCurrencyAndAmountFedwire1SimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCodeFixed                   `xml:"Ccy,attr"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification61 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification62 struct {
	FinInstnId FinancialInstitutionIdentification182 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification63 struct {
	FinInstnId FinancialInstitutionIdentification181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 FinInstnId"`
	BrnchId    *BranchData31                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BrnchId,omitempty"`
}

type BranchData31 struct {
	Id *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	Tp   *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Tp,omitempty"`
	Ccy  *ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Ccy,omitempty"`
	Nm   *Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm,omitempty"`
	Prxy *ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type Charges71 struct {
	Amt ActiveOrHistoricCurrencyAndAmount             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Amt"`
	Agt BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Agt"`
}

type ClearingSystemIdentification2Choice1 struct {
	Cd *ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type ClearingSystemIdentification2Choice2 struct {
	Cd *ExternalClearingSystemIdentification1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type ClearingSystemIdentification3Choice1 struct {
	Cd *ExternalCashClearingSystem1CodeFixed `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type ClearingSystemMemberIdentification21 struct {
	ClrSysId ClearingSystemIdentification2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysId"`
	MmbId    Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 MmbId"`
}

type ClearingSystemMemberIdentification22 struct {
	ClrSysId ClearingSystemIdentification2Choice2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysId"`
	MmbId    RoutingNumberFRS1                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 MmbId"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     fedwire.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BirthDt"`
	PrvcOfBirth *Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CityOfBirth"`
	CtryOfBirth CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtryOfBirth"`
}

type FinancialInstitutionIdentification181 struct {
	BICFI       *BICFIDec2014Identifier               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysMmbId,omitempty"`
	LEI         *LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 LEI,omitempty"`
	Nm          *Max140Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm,omitempty"`
	PstlAdr     *PostalAddress241                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstlAdr,omitempty"`
}

type FinancialInstitutionIdentification182 struct {
	ClrSysMmbId ClearingSystemMemberIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSysMmbId"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Issr,omitempty"`
}

type GroupHeader901 struct {
	MsgId    IMADFedwireFunds1       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 MsgId"`
	CreDtTm  fedwire.ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CreDtTm"`
	NbOfTxs  Max15NumericTextFixed   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 NbOfTxs"`
	SttlmInf SettlementInstruction71 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SttlmInf"`
}

type LocalInstrument2Choice1 struct {
	Prtry *LocalInstrumentFedwireFunds1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type OrganisationIdentification291 struct {
	AnyBIC *AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 LEI,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type OriginalGroupInformation291 struct {
	OrgnlMsgId   Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlMsgId"`
	OrgnlMsgNmId MessageNameIdentificationFRS1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlMsgNmId"`
	OrgnlCreDtTm fedwire.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlCreDtTm"`
}

type OriginalTransactionReference321 struct {
	PmtTpInf PaymentTypeInformation271 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PmtTpInf"`
}

type Party38Choice1 struct {
	OrgId  *OrganisationIdentification291 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgId,omitempty"`
	PrvtId *PersonIdentification131       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvtId,omitempty"`
}

type Party40Choice1 struct {
	Pty *PartyIdentification1351 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Pty,omitempty"`
}

type Party40Choice2 struct {
	Pty *PartyIdentification1352                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Agt,omitempty"`
}

type PartyIdentification1351 struct {
	Nm        *Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm,omitempty"`
	PstlAdr   *PostalAddress242 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstlAdr,omitempty"`
	Id        *Party38Choice1   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id,omitempty"`
	CtryOfRes *CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtryOfRes,omitempty"`
}

type PartyIdentification1352 struct {
	Nm        *Max140Text       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Nm,omitempty"`
	PstlAdr   *PostalAddress241 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstlAdr,omitempty"`
	Id        *Party38Choice1   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id,omitempty"`
	CtryOfRes *CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtryOfRes,omitempty"`
}

type PaymentReturnReason61 struct {
	Orgtr    *PartyIdentification1351 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Orgtr,omitempty"`
	Rsn      ReturnReason5Choice1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Rsn"`
	AddtlInf []*Max105Text            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AddtlInf,omitempty"`
}

type PaymentReturnV10 struct {
	GrpHdr GroupHeader901         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 GrpHdr"`
	TxInf  PaymentTransaction1181 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TxInf"`
}

type PaymentTransaction1181 struct {
	RtrId               *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrId,omitempty"`
	OrgnlGrpInf         OriginalGroupInformation291                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlGrpInf"`
	OrgnlInstrId        *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlInstrId,omitempty"`
	OrgnlEndToEndId     *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlEndToEndId,omitempty"`
	OrgnlTxId           *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlTxId,omitempty"`
	OrgnlUETR           UUIDv4Identifier                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlUETR"`
	OrgnlIntrBkSttlmAmt *ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlIntrBkSttlmAmt,omitempty"`
	OrgnlIntrBkSttlmDt  *fedwire.ISODate                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlIntrBkSttlmDt,omitempty"`
	RtrdIntrBkSttlmAmt  ActiveCurrencyAndAmountFedwire1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrdIntrBkSttlmAmt"`
	IntrBkSttlmDt       fedwire.ISODate                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrBkSttlmDt"`
	RtrdInstdAmt        ActiveOrHistoricCurrencyAndAmount             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrdInstdAmt"`
	XchgRate            *BaseOneRate                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 XchgRate,omitempty"`
	CompstnAmt          *ActiveOrHistoricCurrencyAndAmount            `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CompstnAmt,omitempty"`
	ChrgBr              *ChargeBearerType1Code1                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ChrgBr,omitempty"`
	ChrgsInf            []*Charges71                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ChrgsInf,omitempty"`
	InstgAgt            BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 InstgAgt"`
	InstdAgt            BranchAndFinancialInstitutionIdentification62 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 InstdAgt"`
	RtrChain            TransactionParties81                          `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrChain"`
	RtrRsnInf           PaymentReturnReason61                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 RtrRsnInf"`
	OrgnlTxRef          OriginalTransactionReference321               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 OrgnlTxRef"`
}

type PaymentTypeInformation271 struct {
	LclInstrm LocalInstrument2Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 LclInstrm"`
}

type PersonIdentification131 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type PostalAddress241 struct {
	Dept        *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Dept,omitempty"`
	SubDept     *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SubDept,omitempty"`
	StrtNm      *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 StrtNm,omitempty"`
	BldgNb      *Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BldgNb,omitempty"`
	BldgNm      *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BldgNm,omitempty"`
	Flr         *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Flr,omitempty"`
	PstBx       *Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstBx,omitempty"`
	Room        *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Room,omitempty"`
	PstCd       *Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstCd,omitempty"`
	TwnNm       *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TwnNm,omitempty"`
	TwnLctnNm   *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DstrctNm,omitempty"`
	CtrySubDvsn *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtrySubDvsn,omitempty"`
	Ctry        *CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Ctry,omitempty"`
	AdrLine     []*Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AdrLine,omitempty"`
}

type PostalAddress242 struct {
	Dept        *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Dept,omitempty"`
	SubDept     *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SubDept,omitempty"`
	StrtNm      *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 StrtNm,omitempty"`
	BldgNb      *Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BldgNb,omitempty"`
	BldgNm      *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 BldgNm,omitempty"`
	Flr         *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Flr,omitempty"`
	PstBx       *Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstBx,omitempty"`
	Room        *Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Room,omitempty"`
	PstCd       *Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PstCd,omitempty"`
	TwnNm       Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TwnNm"`
	TwnLctnNm   *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DstrctNm,omitempty"`
	CtrySubDvsn *Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CtrySubDvsn,omitempty"`
	Ctry        CountryCode  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Ctry"`
	AdrLine     []*Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp *ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Tp,omitempty"`
	Id Max2048Text              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    *ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
	Prtry *Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Prtry,omitempty"`
}

type ReturnReason5Choice1 struct {
	Cd *ExternalReturnReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cd,omitempty"`
}

type SettlementInstruction71 struct {
	SttlmMtd SettlementMethod1Code1               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 ClrSys"`
}

type TransactionParties81 struct {
	UltmtDbtr         *Party40Choice1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 UltmtDbtr,omitempty"`
	Dbtr              Party40Choice2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Dbtr"`
	DbtrAcct          *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DbtrAcct,omitempty"`
	InitgPty          *Party40Choice1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 InitgPty,omitempty"`
	DbtrAgt           *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DbtrAgt,omitempty"`
	DbtrAgtAcct       *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 DbtrAgtAcct,omitempty"`
	PrvsInstgAgt1     *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvsInstgAgt1,omitempty"`
	PrvsInstgAgt1Acct *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvsInstgAgt1Acct,omitempty"`
	PrvsInstgAgt2     *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvsInstgAgt2,omitempty"`
	PrvsInstgAgt2Acct *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvsInstgAgt2Acct,omitempty"`
	PrvsInstgAgt3     *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvsInstgAgt3,omitempty"`
	PrvsInstgAgt3Acct *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 PrvsInstgAgt3Acct,omitempty"`
	IntrmyAgt1        *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrmyAgt1,omitempty"`
	IntrmyAgt1Acct    *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrmyAgt1Acct,omitempty"`
	IntrmyAgt2        *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrmyAgt2,omitempty"`
	IntrmyAgt2Acct    *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrmyAgt2Acct,omitempty"`
	IntrmyAgt3        *BranchAndFinancialInstitutionIdentification61 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrmyAgt3,omitempty"`
	IntrmyAgt3Acct    *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 IntrmyAgt3Acct,omitempty"`
	CdtrAgt           *BranchAndFinancialInstitutionIdentification63 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CdtrAgt,omitempty"`
	CdtrAgtAcct       *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CdtrAgtAcct,omitempty"`
	Cdtr              Party40Choice2                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 Cdtr"`
	CdtrAcct          *CashAccount38                                 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 CdtrAcct,omitempty"`
	UltmtCdtr         *Party40Choice1                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.004.001.10 UltmtCdtr,omitempty"`
}

// XSD SimpleType declarations

type ActiveCurrencyAndAmountFedwire1SimpleType float64

type ActiveCurrencyCodeFixed string

const ActiveCurrencyCodeFixedUsd ActiveCurrencyCodeFixed = "USD"

type ActiveOrHistoricCurrencyAndAmountSimpleType fedwire.Amount

type ActiveOrHistoricCurrencyCode string

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type BaseOneRate float64

type ChargeBearerType1Code1 string

const ChargeBearerType1Code1Cred ChargeBearerType1Code1 = "CRED"
const ChargeBearerType1Code1Shar ChargeBearerType1Code1 = "SHAR"
const ChargeBearerType1Code1Slev ChargeBearerType1Code1 = "SLEV"

type CountryCode string

type ExternalAccountIdentification1Code string

type ExternalCashAccountType1Code string

type ExternalCashClearingSystem1CodeFixed string

const ExternalCashClearingSystem1CodeFixedFdw ExternalCashClearingSystem1CodeFixed = "FDW"

type ExternalClearingSystemIdentification1Code string

type ExternalClearingSystemIdentification1CodeFixed string

const ExternalClearingSystemIdentification1CodeFixedUsaba ExternalClearingSystemIdentification1CodeFixed = "USABA"

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type ExternalProxyAccountType1Code string

type ExternalReturnReason1Code string

type IBAN2007Identifier string

type IMADFedwireFunds1 string

type LEIIdentifier string

type LocalInstrumentFedwireFunds1 string

const LocalInstrumentFedwireFunds1Btrc LocalInstrumentFedwireFunds1 = "BTRC"
const LocalInstrumentFedwireFunds1Ctrc LocalInstrumentFedwireFunds1 = "CTRC"
const LocalInstrumentFedwireFunds1Btrs LocalInstrumentFedwireFunds1 = "BTRS"
const LocalInstrumentFedwireFunds1Ctrd LocalInstrumentFedwireFunds1 = "CTRD"
const LocalInstrumentFedwireFunds1Covc LocalInstrumentFedwireFunds1 = "COVC"
const LocalInstrumentFedwireFunds1Covs LocalInstrumentFedwireFunds1 = "COVS"
const LocalInstrumentFedwireFunds1Btrd LocalInstrumentFedwireFunds1 = "BTRD"
const LocalInstrumentFedwireFunds1Ctrs LocalInstrumentFedwireFunds1 = "CTRS"

type Max105Text string

type Max140Text string

type Max15NumericTextFixed string

const Max15NumericTextFixed1 Max15NumericTextFixed = "1"

type Max16Text string

type Max2048Text string

type Max34Text string

type Max35Text string

type Max70Text string

type MessageNameIdentificationFRS1 string

type RoutingNumberFRS1 string

type SettlementMethod1Code1 string

const SettlementMethod1Code1Clrg SettlementMethod1Code1 = "CLRG"

type UUIDv4Identifier string
