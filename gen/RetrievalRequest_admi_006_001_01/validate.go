// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:admi.006.001.01
package RetrievalRequest_admi_006_001_01

import (
	"github.com/moov-io/base"
	"github.com/moov-io/fedwire20022/pkg/fedwire"
)

// XSD Element validations

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	fedwire.AddError(&errs, baseName+".RsndReq", v.RsndReq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations

func (v GenericIdentification11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericIdentification11"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericIdentification361) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericIdentification361"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	fedwire.AddError(&errs, baseName+".Issr", v.Issr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MessageHeader71) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MessageHeader71"
	fedwire.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	fedwire.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	fedwire.AddError(&errs, baseName+".ReqTp", v.ReqTp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification120Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification120Choice1"
	if v.PrtryId != nil {
		fedwire.AddError(&errs, baseName+".PrtryId", v.PrtryId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification1361) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification1361"
	fedwire.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RequestType4Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RequestType4Choice1"
	if v.Prtry != nil {
		fedwire.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ResendRequestV01) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ResendRequestV01"
	fedwire.AddError(&errs, baseName+".MsgHdr", v.MsgHdr.Validate())
	fedwire.AddError(&errs, baseName+".RsndSchCrit", v.RsndSchCrit.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ResendSearchCriteria21) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ResendSearchCriteria21"
	fedwire.AddError(&errs, baseName+".BizDt", v.BizDt.Validate())
	if v.SeqRg != nil {
		fedwire.AddError(&errs, baseName+".SeqRg", v.SeqRg.Validate())
	}
	if v.OrgnlMsgNmId != nil {
		fedwire.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	}
	if v.FileRef != nil {
		fedwire.AddError(&errs, baseName+".FileRef", v.FileRef.Validate())
	}
	fedwire.AddError(&errs, baseName+".Rcpt", v.Rcpt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SequenceRange1Choice1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SequenceRange1Choice1"
	for indx := range v.FrToSeq {
		fedwire.AddError(&errs, baseName+".FrToSeq", v.FrToSeq[indx].Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SequenceRange11) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SequenceRange11"
	fedwire.AddError(&errs, baseName+".FrSeq", v.FrSeq.Validate())
	fedwire.AddError(&errs, baseName+".ToSeq", v.ToSeq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v EndpointIdentifierFedwireFunds1) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[A-Z0-9]{8,8}`); err != nil {
		return err
	}
	if err := fedwire.ValidateMinLength(string(v), 8); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 8); err != nil {
		return err
	}
	return nil
}

func (v IMADOrOMADFedwireFunds1) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[0-9]{8}[A-Z0-9]{8}[0-9]{6}([0-9]{4}[0-9]{4}[A-Z0-9]{4}){0,1}`); err != nil {
		return err
	}
	if err := fedwire.ValidateMinLength(string(v), 22); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxLength(string(v), 34); err != nil {
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

func (v Max35TextFixed) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "NA"); err != nil {
		return err
	}
	return nil
}

func (v MessageNameIdentificationFRS1) Validate() error {
	if err := fedwire.ValidatePattern(string(v), `[a-z]{4,4}[.]{1,1}[0-9]{3,3}[.]{1,1}001[.]{1,1}[0-9]{2,2}`); err != nil {
		return err
	}
	if err := fedwire.ValidateLength(string(v), 15); err != nil {
		return err
	}
	return nil
}

func (v TrafficTypeFedwireFunds1) Validate() error {
	if err := fedwire.ValidateEnumeration(string(v), "S", "R"); err != nil {
		return err
	}
	return nil
}

func (v XSequenceNumberFedwireFunds1) Validate() error {
	if err := fedwire.ValidateMinInclusive(float64(v), 000001); err != nil {
		return err
	}
	if err := fedwire.ValidateMaxInclusive(float64(v), 999999); err != nil {
		return err
	}
	return nil
}
