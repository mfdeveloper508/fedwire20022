// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:camt.060.001.04
package camt_060_001_04

import (
	"github.com/moov-io/base"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Element validations

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	fedwire.AddError(&errs, baseName+".AcctRptgReq", v.AcctRptgReq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations

func (v AccountIdentification4Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountIdentification4Choice"
	if v.IBAN != nil {
		fedwire.AddError(&errs, baseName+".IBAN", v.IBAN.Validate())
	}
	if v.Othr != nil {
		fedwire.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AccountReportingRequestV04) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountReportingRequestV04"
	fedwire.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	for indx := range v.RptgReq {
		fedwire.AddError(&errs, baseName+".RptgReq", v.RptgReq[indx].Validate())
	}
	if v.SplmtryData != nil {
		for indx := range v.SplmtryData {
			fedwire.AddError(&errs, baseName+".SplmtryData", v.SplmtryData[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AccountSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountSchemeName1Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ActiveOrHistoricCurrencyAndAmount) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveOrHistoricCurrencyAndAmount"

	fedwire.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BalanceSubType1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BalanceSubType1Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BalanceType10Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BalanceType10Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BalanceType13) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BalanceType13"
	fedwire.AddError(&errs, baseName+".CdOrPrtry", v.CdOrPrtry.Validate())
	if v.SubTp != nil {
		fedwire.AddError(&errs, baseName+".SubTp", v.SubTp.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification5"
	fedwire.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if v.BrnchId != nil {
		fedwire.AddError(&errs, baseName+".BrnchId", v.BrnchId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchData2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchData2"
	if v.Id != nil {
		fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.Nm != nil {
		fedwire.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		fedwire.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccount24) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccount24"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.Tp != nil {
		fedwire.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	}
	if v.Ccy != nil {
		fedwire.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())
	}
	if v.Nm != nil {
		fedwire.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CashAccountType2Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CashAccountType2Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemIdentification2Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemIdentification2Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification2"
	if v.ClrSysId != nil {
		fedwire.AddError(&errs, baseName+".ClrSysId", v.ClrSysId.Validate())
	}
	fedwire.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ContactDetails2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ContactDetails2"
	if v.NmPrfx != nil {
		fedwire.AddError(&errs, baseName+".NmPrfx", v.NmPrfx.Validate())
	}
	if v.Nm != nil {
		fedwire.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PhneNb != nil {
		fedwire.AddError(&errs, baseName+".PhneNb", v.PhneNb.Validate())
	}
	if v.MobNb != nil {
		fedwire.AddError(&errs, baseName+".MobNb", v.MobNb.Validate())
	}
	if v.FaxNb != nil {
		fedwire.AddError(&errs, baseName+".FaxNb", v.FaxNb.Validate())
	}
	if v.EmailAdr != nil {
		fedwire.AddError(&errs, baseName+".EmailAdr", v.EmailAdr.Validate())
	}
	if v.Othr != nil {
		fedwire.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndPlaceOfBirth1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndPlaceOfBirth1"
	fedwire.AddError(&errs, baseName+".BirthDt", v.BirthDt.Validate())
	if v.PrvcOfBirth != nil {
		fedwire.AddError(&errs, baseName+".PrvcOfBirth", v.PrvcOfBirth.Validate())
	}
	fedwire.AddError(&errs, baseName+".CityOfBirth", v.CityOfBirth.Validate())
	fedwire.AddError(&errs, baseName+".CtryOfBirth", v.CtryOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DatePeriodDetails1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DatePeriodDetails1"
	fedwire.AddError(&errs, baseName+".FrDt", v.FrDt.Validate())
	if v.ToDt != nil {
		fedwire.AddError(&errs, baseName+".ToDt", v.ToDt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v EntryStatus1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "EntryStatus1Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialIdentificationSchemeName1Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification8) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification8"
	if v.BICFI != nil {
		fedwire.AddError(&errs, baseName+".BICFI", v.BICFI.Validate())
	}
	if v.ClrSysMmbId != nil {
		fedwire.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	}
	if v.Nm != nil {
		fedwire.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		fedwire.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Othr != nil {
		fedwire.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericAccountIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericAccountIdentification1"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		fedwire.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		fedwire.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericFinancialIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericFinancialIdentification1"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		fedwire.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		fedwire.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification1"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		fedwire.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		fedwire.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericPersonIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericPersonIdentification1"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		fedwire.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if v.Issr != nil {
		fedwire.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader76) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader76"
	fedwire.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	fedwire.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if v.MsgSndr != nil {
		fedwire.AddError(&errs, baseName+".MsgSndr", v.MsgSndr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Limit2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Limit2"
	fedwire.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	fedwire.AddError(&errs, baseName+".CdtDbtInd", v.CdtDbtInd.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification8) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification8"
	if v.AnyBIC != nil {
		fedwire.AddError(&errs, baseName+".AnyBIC", v.AnyBIC.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			fedwire.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentificationSchemeName1Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party34Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party34Choice"
	if v.OrgId != nil {
		fedwire.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		fedwire.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party35Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party35Choice"
	if v.Pty != nil {
		fedwire.AddError(&errs, baseName+".Pty", v.Pty.Validate())
	}
	if v.Agt != nil {
		fedwire.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification125) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification125"
	if v.Nm != nil {
		fedwire.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		fedwire.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtryOfRes != nil {
		fedwire.AddError(&errs, baseName+".CtryOfRes", v.CtryOfRes.Validate())
	}
	if v.CtctDtls != nil {
		fedwire.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13"
	if v.DtAndPlcOfBirth != nil {
		fedwire.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			fedwire.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentificationSchemeName1Choice"
	if v.Cd != nil {
		fedwire.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress6) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress6"
	if v.AdrTp != nil {
		fedwire.AddError(&errs, baseName+".AdrTp", v.AdrTp.Validate())
	}
	if v.Dept != nil {
		fedwire.AddError(&errs, baseName+".Dept", v.Dept.Validate())
	}
	if v.SubDept != nil {
		fedwire.AddError(&errs, baseName+".SubDept", v.SubDept.Validate())
	}
	if v.StrtNm != nil {
		fedwire.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	}
	if v.BldgNb != nil {
		fedwire.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	if v.PstCd != nil {
		fedwire.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	}
	if v.TwnNm != nil {
		fedwire.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	}
	if v.CtrySubDvsn != nil {
		fedwire.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	}
	if v.Ctry != nil {
		fedwire.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	}
	if v.AdrLine != nil {
		for indx := range v.AdrLine {
			fedwire.AddError(&errs, baseName+".AdrLine", v.AdrLine[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ReportingPeriod2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ReportingPeriod2"
	fedwire.AddError(&errs, baseName+".FrToDt", v.FrToDt.Validate())
	if v.FrToTm != nil {
		fedwire.AddError(&errs, baseName+".FrToTm", v.FrToTm.Validate())
	}
	fedwire.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ReportingRequest4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ReportingRequest4"
	if v.Id != nil {
		fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	fedwire.AddError(&errs, baseName+".ReqdMsgNmId", v.ReqdMsgNmId.Validate())
	if v.Acct != nil {
		fedwire.AddError(&errs, baseName+".Acct", v.Acct.Validate())
	}
	fedwire.AddError(&errs, baseName+".AcctOwnr", v.AcctOwnr.Validate())
	if v.AcctSvcr != nil {
		fedwire.AddError(&errs, baseName+".AcctSvcr", v.AcctSvcr.Validate())
	}
	if v.RptgPrd != nil {
		fedwire.AddError(&errs, baseName+".RptgPrd", v.RptgPrd.Validate())
	}
	if v.RptgSeq != nil {
		fedwire.AddError(&errs, baseName+".RptgSeq", v.RptgSeq.Validate())
	}
	if v.ReqdTxTp != nil {
		fedwire.AddError(&errs, baseName+".ReqdTxTp", v.ReqdTxTp.Validate())
	}
	if v.ReqdBalTp != nil {
		for indx := range v.ReqdBalTp {
			fedwire.AddError(&errs, baseName+".ReqdBalTp", v.ReqdBalTp[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SequenceRange1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SequenceRange1"
	fedwire.AddError(&errs, baseName+".FrSeq", v.FrSeq.Validate())
	fedwire.AddError(&errs, baseName+".ToSeq", v.ToSeq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SequenceRange1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SequenceRange1Choice"
	if v.FrSeq != nil {
		fedwire.AddError(&errs, baseName+".FrSeq", v.FrSeq.Validate())
	}
	if v.ToSeq != nil {
		fedwire.AddError(&errs, baseName+".ToSeq", v.ToSeq.Validate())
	}
	for indx := range v.FrToSeq {
		fedwire.AddError(&errs, baseName+".FrToSeq", v.FrToSeq[indx].Validate())
	}
	for indx := range v.EQSeq {
		fedwire.AddError(&errs, baseName+".EQSeq", v.EQSeq[indx].Validate())
	}
	for indx := range v.NEQSeq {
		fedwire.AddError(&errs, baseName+".NEQSeq", v.NEQSeq[indx].Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SupplementaryData1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SupplementaryData1"
	if v.PlcAndNm != nil {
		fedwire.AddError(&errs, baseName+".PlcAndNm", v.PlcAndNm.Validate())
	}
	fedwire.AddError(&errs, baseName+".Envlp", v.Envlp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SupplementaryDataEnvelope1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TimePeriodDetails1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TimePeriodDetails1"
	fedwire.AddError(&errs, baseName+".FrTm", v.FrTm.Validate())
	if v.ToTm != nil {
		fedwire.AddError(&errs, baseName+".ToTm", v.ToTm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v TransactionType2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "TransactionType2"
	fedwire.AddError(&errs, baseName+".Sts", v.Sts.Validate())
	fedwire.AddError(&errs, baseName+".CdtDbtInd", v.CdtDbtInd.Validate())
	if v.FlrLmt != nil {
		for indx := range v.FlrLmt {
			fedwire.AddError(&errs, baseName+".FlrLmt", v.FlrLmt[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v ActiveOrHistoricCurrencyCode) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z]{3,3}`); err != nil {
		return err
	}
	return nil
}

func (v AddressType2Code) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "ADDR", "PBOX", "HOME", "BIZZ", "MLTO", "DLVY"); err != nil {
		return err
	}
	return nil
}

func (v AnyBICIdentifier) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`); err != nil {
		return err
	}
	return nil
}

func (v BICFIIdentifier) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}`); err != nil {
		return err
	}
	return nil
}

func (v CountryCode) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z]{2,2}`); err != nil {
		return err
	}
	return nil
}

func (v CreditDebitCode) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "CRDT", "DBIT"); err != nil {
		return err
	}
	return nil
}

func (v ExternalAccountIdentification1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalBalanceSubType1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalBalanceType1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalCashAccountType1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalClearingSystemIdentification1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 5); err != nil {
		return err
	}
	return nil
}

func (v ExternalEntryStatus1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalFinancialInstitutionIdentification1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalOrganisationIdentification1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v ExternalPersonIdentification1Code) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v FloorLimitType1Code) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "CRED", "DEBT", "BOTH"); err != nil {
		return err
	}
	return nil
}

func (v IBAN2007Identifier) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v ISOTime) Validate() error {
	return nil
}

func (v Max140Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 140); err != nil {
		return err
	}
	return nil
}

func (v Max16Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 16); err != nil {
		return err
	}
	return nil
}

func (v Max2048Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 2048); err != nil {
		return err
	}
	return nil
}

func (v Max34Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 34); err != nil {
		return err
	}
	return nil
}

func (v Max350Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 350); err != nil {
		return err
	}
	return nil
}

func (v Max35Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max70Text) Validate() error {
	if err := fedwire.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 70); err != nil {
		return err
	}
	return nil
}

func (v NamePrefix1Code) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "DOCT", "MIST", "MISS", "MADM"); err != nil {
		return err
	}
	return nil
}

func (v PhoneNumber) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `\+[0-9]{1,3}-[0-9()+\-]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v QueryType3Code) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "ALLL", "CHNG", "MODF"); err != nil {
		return err
	}
	return nil
}
