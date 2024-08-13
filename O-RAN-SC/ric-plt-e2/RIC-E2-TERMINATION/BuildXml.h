/*
 * Copyright 2020 AT&T Intellectual Property
 * Copyright 2020 Nokia
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//
// Created by adi ENZEL on 4/5/20.
//

#ifndef E2_BUILDXML_H
#define E2_BUILDXML_H

#include <iostream>
#include <iosfwd>
#include <vector>
#include "pugixml/src/pugixml.hpp"
#include <string>
#include <sstream>
#include <mdclog/mdclog.h>
#include <cstdlib>

using namespace std;

/*
 * Copied from pugixml samples
 */
struct xml_string_writer : pugi::xml_writer {
    std::string result;

    void write(const void *data, size_t size) override {
        result.append(static_cast<const char *>(data), size);
    }
};

std::string node_to_string(pugi::xml_node node) {
    xml_string_writer writer;
    node.print(writer);

    return writer.result;
}

string buildXmlData(const string &messageName,
        const string &ieName,
        vector<string> &RANfunctionsAdded,
        vector<string> &RANfunctionsModified,
        unsigned char *buffer,
        size_t size) {
    pugi::xml_document *doc = new pugi::xml_document();

    char RANfunctionsAddedID[8];
    snprintf(RANfunctionsAddedID, 8, "%d", (int) ProtocolIE_ID_id_RANfunctionsAdded);
    char RANfunctionsModifiedID[8];
    snprintf(RANfunctionsModifiedID, 8, "%d", (int) ProtocolIE_ID_id_RANfunctionsModified);
    char GlobalE2nodeID[8];
    snprintf(GlobalE2nodeID, 8, "%d", (int) ProtocolIE_ID_id_GlobalE2node_ID);

    pugi::xml_parse_result result = doc->load_buffer((const char *) buffer, size);
    if (result) {
        for (auto tool : doc->child("E2AP-PDU")
                .child("initiatingMessage")
                .child("value")
                .child(messageName.c_str())
                .child("protocolIEs")
                .children(ieName.c_str())) {
            // there can be many ieName entries in the messageName so we need only the ones that containes E2SM continers
            auto node = tool.child("id");  // get the id to identify the type of the contained message
            if (node.empty()) {
                mdclog_write(MDCLOG_ERR, "Failed to find ID node in the XML. File %s, line %d",
                             __FILE__, __LINE__);
                continue;
            }
            if (strcmp(node.name(), "id") == 0 && strcmp(node.child_value(), RANfunctionsAddedID) == 0) {
                auto nodea = tool.child("value").
                        child("RANfunctions-List").
                        children("ProtocolIE-SingleContainer");
                unsigned int index = 0;
                for (auto n1 : nodea) {
                    auto n2 = n1.child("value").child("RANfunction-Item").child("ranFunctionDefinition");
                    n2.remove_children();
                    string val = RANfunctionsAdded.at(index++);
                    // here we get vector with counter
                    n2.append_child(pugi::node_pcdata).set_value(val.c_str());
                    if (mdclog_level_get() >= MDCLOG_DEBUG) {
                        mdclog_write(MDCLOG_DEBUG, "entry %s Replaced with : %s", n2.name(), n2.child_value());
                    }
                }
            } else if (strcmp(node.name(), "id") == 0 && strcmp(node.child_value(), RANfunctionsModifiedID) == 0) {
                auto nodea = tool.child("value").
                        child("RANfunctions-List").
                        children("ProtocolIE-SingleContainer");
                unsigned int index = 0;
                for (auto n1 : nodea) {
                    auto n2 = n1.child("value").child("RANfunction-Item").child("ranFunctionDefinition");
                    n2.remove_children();
                    string val = RANfunctionsModified.at(index++);
                    // here we get vector with counter
                    n2.append_child(pugi::node_pcdata).set_value(val.c_str());
                    if (mdclog_level_get() >= MDCLOG_DEBUG) {
                        mdclog_write(MDCLOG_DEBUG, "entry %s Replaced with : %s", n2.name(), n2.child_value());
                    }
                }
            } else if (strcmp(node.name(), "id") == 0 && strcmp(node.child_value(), GlobalE2nodeID) == 0) {
                continue;
            } else {
                if (mdclog_level_get() >= MDCLOG_DEBUG) {
                    mdclog_write(MDCLOG_DEBUG, "Entry name :%s  with entry of %s skipped", node.name(), node.child_value());
                }
                continue;
            }
        }
    } else {
        mdclog_write(MDCLOG_ERR,
                     "Error loading xml string");
        delete doc;
        return string("");
    }
    auto res = node_to_string(*doc);
    res.erase(std::remove(res.begin(), res.end(), '\n'), res.end());
    res.erase(std::remove(res.begin(), res.end(), '\t'), res.end());
    delete doc;
    return res;

}

#endif //E2_BUILDXML_H
