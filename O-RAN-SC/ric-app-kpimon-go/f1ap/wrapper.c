#include"wrapper.h"
// Function to decode a single Base64 character
unsigned char decodeBase64Char(char c) {
   if (c >= 'A' && c <= 'Z') return c - 'A';
   if (c >= 'a' && c <= 'z') return c - 'a' + 26;
   if (c >= '0' && c <= '9') return c - '0' + 52;
   if (c == '+') return 62;
   if (c == '/') return 63;
   return 0;  // Invalid character
}
// Function to decode a Base64 encoded string
unsigned char* base64Decode(const char* input, size_t* outputLength) {
   size_t inputLength = strlen(input);
   if (inputLength % 4 != 0) {
       printf("Invalid Base64 input length.\n");
       return NULL;
   }

   *outputLength = (inputLength / 4) * 3;
   unsigned char* output = (unsigned char*)malloc(*outputLength);
   if (output == NULL) {
       printf("Memory allocation failed.\n");
       return NULL;
   }

   for (size_t i = 0, j = 0; i < inputLength; i += 4, j += 3) {
       unsigned char a = decodeBase64Char(input[i]);
       unsigned char b = decodeBase64Char(input[i + 1]);
       unsigned char c = decodeBase64Char(input[i + 2]);
       unsigned char d = decodeBase64Char(input[i + 3]);

       output[j] = (a << 2) | (b >> 4);
       output[j + 1] = (b << 4) | (c >> 2);
       output[j + 2] = (c << 6) | d;
   }

   return output;
}
					  ///encodedString
CellIds_t decodeF1apGetCellIds(const char* decoded ){   
	/*
	   size_t decodedLength;
   unsigned char* decoded = base64Decode(encodedString, &decodedLength);
   */
	// printf("ZONGA\n");
	CellIds_t res;
   if (decoded != NULL) {
       printf("Decoded string: %s\n", decoded);
       //free(decoded);
   }
       // Calculate the length of the hex string
        size_t hex_len = strlen(decoded);

        // Allocate memory for a char array to store the hex values
        char *hex_buffer = (char *)malloc(hex_len / 2 + 1); // Each byte is represented by 2 characters, +1 for null terminator

        if (hex_buffer == NULL) {
                fprintf(stderr, "Memory allocation failed\n");
                return res;
        }

        // Convert the hex string to binary data
        for (size_t i = 0; i < hex_len; i += 2) {
                char byte[3] = {decoded[i], decoded[i + 1], '\0'};
                hex_buffer[i / 2] = (char)strtol(byte, NULL, 16);
        }

        // Null-terminate the char array
        hex_buffer[hex_len / 2] = '\0';

        // Now hex_buffer contains the binary data corresponding to the hex values

        // Print the result
        printf("Hex values as a string: %s\n", hex_buffer);

        F1AP_PDU_t *f1pdu=(F1AP_PDU_t *)calloc(1,sizeof(F1AP_PDU_t ));
        char** cellList;
	char** plmnList;
        int cellListLength=0;
        enum asn_transfer_syntax syntax;

        syntax = ATS_ALIGNED_BASIC_PER;
	printf("F1AP\n");
        asn_dec_rval_t rval =  asn_decode(NULL, syntax, &asn_DEF_F1AP_PDU, (void**)&f1pdu, hex_buffer, hex_len);
        typedef Served_Cells_To_Add_ItemIEs_t    ProtocolIE_SingleContainer;

             if(rval.code == RC_OK)
        {
                printf( "[INFO] F1AP PDU decoded successfull rval.code = %d \n",rval.code);

                //asn_fprint(stdout, &asn_DEF_F1AP_PDU, f1pdu);
                xer_fprint(stdout, &asn_DEF_F1AP_PDU, f1pdu);
                //printf("present value =: %d\n", f1pdu->present);
                if( f1pdu->present==1)//initiating message
                {
                        if(f1pdu->choice.initiatingMessage->value.present==3)//GNBDUConfigurationUpdate
                        {
                                int count=f1pdu->choice.initiatingMessage->value.choice.GNBDUConfigurationUpdate.protocolIEs.list.count;
                                 //printf("count value =: %d\n", count);
                                 for(int i=0; i<count;i++){
                                        GNBDUConfigurationUpdateIEs_t * tmpVar= f1pdu->choice.initiatingMessage->value.choice.GNBDUConfigurationUpdate.protocolIEs.list.array[i];
                                        if(tmpVar->value.present==2)//Served_Cells_To_Add_List
                                        {
                                                cellListLength=tmpVar->value.choice.Served_Cells_To_Add_List.list.count;
                                                printf("cellListLength value =: %d\n", cellListLength);
                                                cellList= (char **)malloc(cellListLength*sizeof(char *));
						plmnList= (char **)malloc(cellListLength*sizeof(char *));
                                                for(int j=0;j<cellListLength;j++ )
                                                {
                                                        if(tmpVar->value.choice.Served_Cells_To_Add_List.list.array[j]->value.present==1)//Served_Cells_To_Add_Item

                                                        {
                                                                NRCGI_t NRC=tmpVar->value.choice.Served_Cells_To_Add_List.list.array[j]->value.choice.Served_Cells_To_Add_Item.served_Cell_Information.nRCGI;
                                                                //xer_fprint(stdout, &asn_DEF_NRCGI, &NRC);
								/*
                                                                std::stringstream pLMN_HexStringStream;
                                                                for (int i = 0; i < NRC.pLMN_Identity.size; ++i)
                                                                {
                                                                        pLMN_HexStringStream << std::setw(2) << std::setfill('0') << std::hex << static_cast<int>(NRC.pLMN_Identity.buf[i]);

                                                                }
                                                                std::stringstream nRCell_HexStringStream;
                                                                for (int i = 0; i < NRC.nRCellIdentity.size; ++i)
                                                                {
                                                                        nRCell_HexStringStream << std::setw(2) << std::setfill('0') << std::hex << static_cast<int>(NRC.nRCellIdentity.buf[i]);
                                                                }
                                                                std::string nRCell_HexString = nRCell_HexStringStream.str();
                                                                std::string pLMN_HexString = pLMN_HexStringStream.str();
                                                                for(int k=0; k<pLMN_HexString.size(); k++){
                                                                        pLMN_HexString[k]=toupper(pLMN_HexString[k]);
                                                                }
                                                                for(int k=0; k<nRCell_HexString.size(); k++){
                                                                        nRCell_HexString[k]=toupper(nRCell_HexString[k]);
                                                                }
								*/
								char ans[2*NRC.pLMN_Identity.size];
                                                                int ans_index=0;
                                                                 for (size_t i = 0; i < NRC.pLMN_Identity.size; i++) {
                                                                        ans_index += sprintf(&ans[ans_index], "%02X", NRC.pLMN_Identity.buf[i]);
                                                                }
                                                                ans[ans_index] = '\0'; // Null-terminate the string
                                                                // Print the hexadecimal string
                                                                printf("PLMN-Identity as hexadecimal string: %s\n", ans);
								plmnList[j]=(char*)malloc(2*NRC.pLMN_Identity.size);
								strcpy(plmnList[j], ans);
                                                                int ans_index2=0;
                                                                char ans2[(NRC.nRCellIdentity.size * 2) + 1]; // Make sure ans has enough space for the binary string
                                                                for (size_t i = 0; i < NRC.nRCellIdentity.size; i++) {
                                                                        ans_index2 += sprintf(&ans2[ans_index2], "%02X",NRC.nRCellIdentity.buf[i]);
                                                                }
                                                                ans2[ans_index2] = '\0'; // Null-terminate the string

                                                                // Print the hexadecimal string
                                                                printf("Bit String as hexadecimal string: %s\n", ans2);


                                                                cellList[j]=(char*)malloc((NRC.nRCellIdentity.size * 2) + 1);
                                                                strcpy(cellList[j], ans2);

                                                                //std::cout << "pLMN-Identity as hexadecimal string: "<<pLMN_HexString<< std::endl;
                                                                 //std::cout << "nrrcellID as hexadecimal string: " <<nRCell_HexString<< std::endl;


                                                        }
                                                }

                                        }
                                 }

                        }
                }
        }

	     else{
	     	printf("F1AP decoding Failed\n");
	     }
	     /*
             for(int i=0;i<cellListLength;i++){
                std::cout<<cellList[i]<<" ";
             }
	     */
	     res.cellids=cellList;
	     res.size=cellListLength;
	     res.plmn=plmnList;
	     return res;

}

