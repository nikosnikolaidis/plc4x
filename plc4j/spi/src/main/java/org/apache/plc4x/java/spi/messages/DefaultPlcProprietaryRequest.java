/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.messages;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.PlcResponse;

import java.util.concurrent.CompletableFuture;

@JsonTypeInfo(use = JsonTypeInfo.Id.CLASS, property = "className")
public class DefaultPlcProprietaryRequest<REQUEST> implements InternalPlcProprietaryRequest<REQUEST> {

    @Override
    @JsonIgnore
    public CompletableFuture<PlcResponse> execute() {
        throw new PlcRuntimeException("not supported"); // TODO: figure out what to do with this
    }

    private REQUEST proprietaryRequest;

    @JsonCreator(mode = JsonCreator.Mode.PROPERTIES)
    public DefaultPlcProprietaryRequest(@JsonProperty("proprietaryRequest") REQUEST proprietaryRequest) {
        this.proprietaryRequest = proprietaryRequest;
    }

    @Override
    public REQUEST getProprietaryRequest() {
        return proprietaryRequest;
    }
}