//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "encoding/xml"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetConfirmedServiceACKGetEnrollmentSummary struct {
    BACnetConfirmedServiceACK
}

// The corresponding interface
type IBACnetConfirmedServiceACKGetEnrollmentSummary interface {
    IBACnetConfirmedServiceACK
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetConfirmedServiceACKGetEnrollmentSummary) ServiceChoice() uint8 {
    return 0x04
}

func (m BACnetConfirmedServiceACKGetEnrollmentSummary) initialize() spi.Message {
    return m
}

func NewBACnetConfirmedServiceACKGetEnrollmentSummary() BACnetConfirmedServiceACKInitializer {
    return &BACnetConfirmedServiceACKGetEnrollmentSummary{}
}

func CastIBACnetConfirmedServiceACKGetEnrollmentSummary(structType interface{}) IBACnetConfirmedServiceACKGetEnrollmentSummary {
    castFunc := func(typ interface{}) IBACnetConfirmedServiceACKGetEnrollmentSummary {
        if iBACnetConfirmedServiceACKGetEnrollmentSummary, ok := typ.(IBACnetConfirmedServiceACKGetEnrollmentSummary); ok {
            return iBACnetConfirmedServiceACKGetEnrollmentSummary
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetConfirmedServiceACKGetEnrollmentSummary(structType interface{}) BACnetConfirmedServiceACKGetEnrollmentSummary {
    castFunc := func(typ interface{}) BACnetConfirmedServiceACKGetEnrollmentSummary {
        if sBACnetConfirmedServiceACKGetEnrollmentSummary, ok := typ.(BACnetConfirmedServiceACKGetEnrollmentSummary); ok {
            return sBACnetConfirmedServiceACKGetEnrollmentSummary
        }
        if sBACnetConfirmedServiceACKGetEnrollmentSummary, ok := typ.(*BACnetConfirmedServiceACKGetEnrollmentSummary); ok {
            return *sBACnetConfirmedServiceACKGetEnrollmentSummary
        }
        return BACnetConfirmedServiceACKGetEnrollmentSummary{}
    }
    return castFunc(structType)
}

func (m BACnetConfirmedServiceACKGetEnrollmentSummary) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetConfirmedServiceACK.LengthInBits()

    return lengthInBits
}

func (m BACnetConfirmedServiceACKGetEnrollmentSummary) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetConfirmedServiceACKGetEnrollmentSummaryParse(io *utils.ReadBuffer) (BACnetConfirmedServiceACKInitializer, error) {

    // Create the instance
    return NewBACnetConfirmedServiceACKGetEnrollmentSummary(), nil
}

func (m BACnetConfirmedServiceACKGetEnrollmentSummary) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetConfirmedServiceACKSerialize(io, m.BACnetConfirmedServiceACK, CastIBACnetConfirmedServiceACK(m), ser)
}

func (m *BACnetConfirmedServiceACKGetEnrollmentSummary) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
    }
}

func (m BACnetConfirmedServiceACKGetEnrollmentSummary) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetConfirmedServiceACKGetEnrollmentSummary"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

