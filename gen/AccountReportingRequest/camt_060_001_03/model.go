// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:camt.060.001.03
package camt_060_001_03

import (
	"encoding/xml"
	"time"

	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Elements

type Document struct {
	XMLName xml.Name

	AcctRptgReq AccountReportingRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 AcctRptgReq"`
}

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	IBAN *IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 IBAN,omitempty"`
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Othr,omitempty"`
}

type AccountReportingRequestV03 struct {
	GrpHdr      GroupHeader59         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 GrpHdr"`
	RptgReq     []ReportingRequest3   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 RptgReq"`
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SplmtryData,omitempty"`
}

type AccountSchemeName1Choice struct {
	Cd    *ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type ActiveOrHistoricCurrencyAndAmount struct {
	Value ActiveOrHistoricCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode                `xml:"Ccy,attr"`
}

type BalanceSubType1Choice struct {
	Cd    *ExternalBalanceSubType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type BalanceType12 struct {
	CdOrPrtry BalanceType5Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CdOrPrtry"`
	SubTp     *BalanceSubType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SubTp,omitempty"`
}

type BalanceType5Choice struct {
	Cd    *BalanceType12Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FinInstnId"`
	BrnchId    *BranchData2                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      *Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id,omitempty"`
	Nm      *Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Nm,omitempty"`
	PstlAdr *PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PstlAdr,omitempty"`
}

type CashAccount24 struct {
	Id  AccountIdentification4Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id"`
	Tp  *CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Tp,omitempty"`
	Ccy *ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Ccy,omitempty"`
	Nm  *Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Nm,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    *ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    *ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ClrSysId,omitempty"`
	MmbId    Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   *NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 NmPrfx,omitempty"`
	Nm       *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Nm,omitempty"`
	PhneNb   *PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PhneNb,omitempty"`
	MobNb    *PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 MobNb,omitempty"`
	FaxNb    *PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FaxNb,omitempty"`
	EmailAdr *Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 EmailAdr,omitempty"`
	Othr     *Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Othr,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     fedwire.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 BirthDt"`
	PrvcOfBirth *Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CityOfBirth"`
	CtryOfBirth CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CtryOfBirth"`
}

type DatePeriodDetails1 struct {
	FrDt fedwire.ISODate  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FrDt"`
	ToDt *fedwire.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ToDt,omitempty"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    *ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ClrSysMmbId,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id"`
	SchmeNm *AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SchmeNm,omitempty"`
	Issr    *Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SchmeNm,omitempty"`
	Issr    *Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Issr,omitempty"`
}

type GroupHeader59 struct {
	MsgId   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 MsgId"`
	CreDtTm fedwire.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CreDtTm"`
	MsgSndr *Party12Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 MsgSndr,omitempty"`
}

type Limit2 struct {
	Amt       ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Amt"`
	CdtDbtInd FloorLimitType1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CdtDbtInd"`
}

type OrganisationIdentification8 struct {
	AnyBIC *AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 AnyBIC,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type Party11Choice struct {
	OrgId  *OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 OrgId,omitempty"`
	PrvtId *PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PrvtId,omitempty"`
}

type Party12Choice struct {
	Pty *PartyIdentification43                        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Pty,omitempty"`
	Agt *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Agt,omitempty"`
}

type PartyIdentification43 struct {
	Nm        *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Nm,omitempty"`
	PstlAdr   *PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PstlAdr,omitempty"`
	Id        *Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id,omitempty"`
	CtryOfRes *CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CtryOfRes,omitempty"`
	CtctDtls  *ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Prtry,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       *AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 AdrTp,omitempty"`
	Dept        *Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Dept,omitempty"`
	SubDept     *Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 SubDept,omitempty"`
	StrtNm      *Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 StrtNm,omitempty"`
	BldgNb      *Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 BldgNb,omitempty"`
	PstCd       *Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PstCd,omitempty"`
	TwnNm       *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 TwnNm,omitempty"`
	CtrySubDvsn *Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CtrySubDvsn,omitempty"`
	Ctry        *CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Ctry,omitempty"`
	AdrLine     []*Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 AdrLine,omitempty"`
}

type ReportingPeriod1 struct {
	FrToDt DatePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FrToDt"`
	FrToTm TimePeriodDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FrToTm"`
	Tp     QueryType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Tp"`
}

type ReportingRequest3 struct {
	Id          *Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Id,omitempty"`
	ReqdMsgNmId Max35Text                                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ReqdMsgNmId"`
	Acct        *CashAccount24                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Acct,omitempty"`
	AcctOwnr    Party12Choice                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 AcctOwnr"`
	AcctSvcr    *BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 AcctSvcr,omitempty"`
	RptgPrd     *ReportingPeriod1                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 RptgPrd,omitempty"`
	ReqdTxTp    *TransactionType1                             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ReqdTxTp,omitempty"`
	ReqdBalTp   []*BalanceType12                              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ReqdBalTp,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm *Max350Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TimePeriodDetails1 struct {
	FrTm ISOTime  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FrTm"`
	ToTm *ISOTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 ToTm,omitempty"`
}

type TransactionType1 struct {
	Sts       EntryStatus2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 Sts"`
	CdtDbtInd CreditDebitCode  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 CdtDbtInd"`
	FlrLmt    []*Limit2        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.060.001.03 FlrLmt,omitempty"`
}

// XSD SimpleType declarations

type ActiveOrHistoricCurrencyAndAmountSimpleType fedwire.Amount

type ActiveOrHistoricCurrencyCode string

type AddressType2Code string

const AddressType2CodeAddr AddressType2Code = "ADDR"
const AddressType2CodePbox AddressType2Code = "PBOX"
const AddressType2CodeHome AddressType2Code = "HOME"
const AddressType2CodeBizz AddressType2Code = "BIZZ"
const AddressType2CodeMlto AddressType2Code = "MLTO"
const AddressType2CodeDlvy AddressType2Code = "DLVY"

type AnyBICIdentifier string

type BICFIIdentifier string

type BalanceType12Code string

const BalanceType12CodeXpcd BalanceType12Code = "XPCD"
const BalanceType12CodeOpav BalanceType12Code = "OPAV"
const BalanceType12CodeItav BalanceType12Code = "ITAV"
const BalanceType12CodeClav BalanceType12Code = "CLAV"
const BalanceType12CodeFwav BalanceType12Code = "FWAV"
const BalanceType12CodeClbd BalanceType12Code = "CLBD"
const BalanceType12CodeItbd BalanceType12Code = "ITBD"
const BalanceType12CodeOpbd BalanceType12Code = "OPBD"
const BalanceType12CodePrcd BalanceType12Code = "PRCD"
const BalanceType12CodeInfo BalanceType12Code = "INFO"

type CountryCode string

type CreditDebitCode string

const CreditDebitCodeCrdt CreditDebitCode = "CRDT"
const CreditDebitCodeDbit CreditDebitCode = "DBIT"

type EntryStatus2Code string

const EntryStatus2CodeBook EntryStatus2Code = "BOOK"
const EntryStatus2CodePdng EntryStatus2Code = "PDNG"
const EntryStatus2CodeInfo EntryStatus2Code = "INFO"

type ExternalAccountIdentification1Code string

type ExternalBalanceSubType1Code string

type ExternalCashAccountType1Code string

type ExternalClearingSystemIdentification1Code string

type ExternalFinancialInstitutionIdentification1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type FloorLimitType1Code string

const FloorLimitType1CodeCred FloorLimitType1Code = "CRED"
const FloorLimitType1CodeDebt FloorLimitType1Code = "DEBT"
const FloorLimitType1CodeBoth FloorLimitType1Code = "BOTH"

type IBAN2007Identifier string

type ISOTime time.Time

type Max140Text string

type Max16Text string

type Max2048Text string

type Max34Text string

type Max350Text string

type Max35Text string

type Max70Text string

type NamePrefix1Code string

const NamePrefix1CodeDoct NamePrefix1Code = "DOCT"
const NamePrefix1CodeMist NamePrefix1Code = "MIST"
const NamePrefix1CodeMiss NamePrefix1Code = "MISS"
const NamePrefix1CodeMadm NamePrefix1Code = "MADM"

type PhoneNumber string

type QueryType3Code string

const QueryType3CodeAlll QueryType3Code = "ALLL"
const QueryType3CodeChng QueryType3Code = "CHNG"
const QueryType3CodeModf QueryType3Code = "MODF"
