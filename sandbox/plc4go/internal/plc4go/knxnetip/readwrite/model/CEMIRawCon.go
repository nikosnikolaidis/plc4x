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
type CEMIRawCon struct {
    CEMI
}

// The corresponding interface
type ICEMIRawCon interface {
    ICEMI
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m CEMIRawCon) MessageCode() uint8 {
    return 0x2F
}

func (m CEMIRawCon) initialize() spi.Message {
    return m
}

func NewCEMIRawCon() CEMIInitializer {
    return &CEMIRawCon{}
}

func CastICEMIRawCon(structType interface{}) ICEMIRawCon {
    castFunc := func(typ interface{}) ICEMIRawCon {
        if iCEMIRawCon, ok := typ.(ICEMIRawCon); ok {
            return iCEMIRawCon
        }
        return nil
    }
    return castFunc(structType)
}

func CastCEMIRawCon(structType interface{}) CEMIRawCon {
    castFunc := func(typ interface{}) CEMIRawCon {
        if sCEMIRawCon, ok := typ.(CEMIRawCon); ok {
            return sCEMIRawCon
        }
        if sCEMIRawCon, ok := typ.(*CEMIRawCon); ok {
            return *sCEMIRawCon
        }
        return CEMIRawCon{}
    }
    return castFunc(structType)
}

func (m CEMIRawCon) LengthInBits() uint16 {
    var lengthInBits uint16 = m.CEMI.LengthInBits()

    return lengthInBits
}

func (m CEMIRawCon) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIRawConParse(io *utils.ReadBuffer) (CEMIInitializer, error) {

    // Create the instance
    return NewCEMIRawCon(), nil
}

func (m CEMIRawCon) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return CEMISerialize(io, m.CEMI, CastICEMI(m), ser)
}

func (m *CEMIRawCon) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m CEMIRawCon) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.CEMIRawCon"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

