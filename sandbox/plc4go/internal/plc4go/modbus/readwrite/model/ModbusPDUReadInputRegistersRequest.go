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
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type ModbusPDUReadInputRegistersRequest struct {
    StartingAddress uint16
    Quantity uint16
    ModbusPDU
}

// The corresponding interface
type IModbusPDUReadInputRegistersRequest interface {
    IModbusPDU
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m ModbusPDUReadInputRegistersRequest) ErrorFlag() bool {
    return false
}

func (m ModbusPDUReadInputRegistersRequest) FunctionFlag() uint8 {
    return 0x04
}

func (m ModbusPDUReadInputRegistersRequest) Response() bool {
    return false
}

func (m ModbusPDUReadInputRegistersRequest) initialize() spi.Message {
    return m
}

func NewModbusPDUReadInputRegistersRequest(startingAddress uint16, quantity uint16) ModbusPDUInitializer {
    return &ModbusPDUReadInputRegistersRequest{StartingAddress: startingAddress, Quantity: quantity}
}

func CastIModbusPDUReadInputRegistersRequest(structType interface{}) IModbusPDUReadInputRegistersRequest {
    castFunc := func(typ interface{}) IModbusPDUReadInputRegistersRequest {
        if iModbusPDUReadInputRegistersRequest, ok := typ.(IModbusPDUReadInputRegistersRequest); ok {
            return iModbusPDUReadInputRegistersRequest
        }
        return nil
    }
    return castFunc(structType)
}

func CastModbusPDUReadInputRegistersRequest(structType interface{}) ModbusPDUReadInputRegistersRequest {
    castFunc := func(typ interface{}) ModbusPDUReadInputRegistersRequest {
        if sModbusPDUReadInputRegistersRequest, ok := typ.(ModbusPDUReadInputRegistersRequest); ok {
            return sModbusPDUReadInputRegistersRequest
        }
        if sModbusPDUReadInputRegistersRequest, ok := typ.(*ModbusPDUReadInputRegistersRequest); ok {
            return *sModbusPDUReadInputRegistersRequest
        }
        return ModbusPDUReadInputRegistersRequest{}
    }
    return castFunc(structType)
}

func (m ModbusPDUReadInputRegistersRequest) LengthInBits() uint16 {
    var lengthInBits uint16 = m.ModbusPDU.LengthInBits()

    // Simple field (startingAddress)
    lengthInBits += 16

    // Simple field (quantity)
    lengthInBits += 16

    return lengthInBits
}

func (m ModbusPDUReadInputRegistersRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUReadInputRegistersRequestParse(io *utils.ReadBuffer) (ModbusPDUInitializer, error) {

    // Simple Field (startingAddress)
    startingAddress, _startingAddressErr := io.ReadUint16(16)
    if _startingAddressErr != nil {
        return nil, errors.New("Error parsing 'startingAddress' field " + _startingAddressErr.Error())
    }

    // Simple Field (quantity)
    quantity, _quantityErr := io.ReadUint16(16)
    if _quantityErr != nil {
        return nil, errors.New("Error parsing 'quantity' field " + _quantityErr.Error())
    }

    // Create the instance
    return NewModbusPDUReadInputRegistersRequest(startingAddress, quantity), nil
}

func (m ModbusPDUReadInputRegistersRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (startingAddress)
    startingAddress := uint16(m.StartingAddress)
    _startingAddressErr := io.WriteUint16(16, (startingAddress))
    if _startingAddressErr != nil {
        return errors.New("Error serializing 'startingAddress' field " + _startingAddressErr.Error())
    }

    // Simple Field (quantity)
    quantity := uint16(m.Quantity)
    _quantityErr := io.WriteUint16(16, (quantity))
    if _quantityErr != nil {
        return errors.New("Error serializing 'quantity' field " + _quantityErr.Error())
    }

        return nil
    }
    return ModbusPDUSerialize(io, m.ModbusPDU, CastIModbusPDU(m), ser)
}

func (m *ModbusPDUReadInputRegistersRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "startingAddress":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.StartingAddress = data
            case "quantity":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Quantity = data
            }
        }
    }
}

func (m ModbusPDUReadInputRegistersRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadInputRegistersRequest"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.StartingAddress, xml.StartElement{Name: xml.Name{Local: "startingAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Quantity, xml.StartElement{Name: xml.Name{Local: "quantity"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

