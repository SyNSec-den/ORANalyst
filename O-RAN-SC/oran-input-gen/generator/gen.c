#include "gen.h"

int gen_hdr(E2SM_KPM_IndicationHeader_t** output) {
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationHeader;
    void** struct_ptr = (void**)output;
    return asn_random_fill(type_descriptor, struct_ptr, 100000);
}

int mutate_hdr(E2SM_KPM_IndicationHeader_t** output) {
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationHeader;
    void** struct_ptr = (void**)output;
    return asn_mutate(type_descriptor, struct_ptr);
}

ssize_t encode_hdr(E2SM_KPM_IndicationHeader_t* input, void* buffer) {
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationHeader;
    void* struct_ptr = (void*)input;
    return aper_encode_to_new_buffer(type_descriptor, 0, struct_ptr, &buffer);
}

asn_dec_rval_t decode_hdr_file(const char* filename, E2SM_KPM_IndicationHeader_t** output) {
    // Open the file
    FILE *file = fopen(filename, "rb");
    if (!file) {
        fprintf(stderr, "Failed to open file: %s\n", filename);
        exit(1);
    }

    // Get the size of the file
    fseek(file, 0, SEEK_END);
    long size = ftell(file);
    fseek(file, 0, SEEK_SET);

    // Read the file contents into a buffer
    char* buffer = (char*)malloc(size);
    if (!fread(buffer, size, 1, file)) {
        fprintf(stderr, "Failed to read file: %s\n", filename);
        exit(1);
    }

    fclose(file);

    // Prepare for decoding
    const asn_codec_ctx_t* opt_codec_ctx = NULL;  
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationHeader;  
    void** struct_ptr = (void**)output;
    const void* data_buffer = (void*)buffer;
    int skip_bits = 0;
    int unused_bits = 0;

    // Call aper_decode
    asn_dec_rval_t rval = aper_decode(opt_codec_ctx, type_descriptor, struct_ptr, data_buffer, size, 0, 0);
    return rval;
}

int mutateHeaderTest() {
    E2SM_KPM_IndicationHeader_t* output = NULL;
    int res;
    // res = gen_hdr(&output);
    // if (res == 0) {
    //     std::cout << "Random gen successful\n";
    //     E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationHeader,output);
    //     std::cout << "\n";
    // } else {
    //     std::cerr << "Random gen failed\n";
    // }

    asn_dec_rval_t eval = decode_hdr_file("/home/tianchang/Desktop/asn1_test/bin_out/header_1696603955_1088799629.dat", &output);
    if (eval.code == RC_OK) {
        printf("Decoding successful\n");
        E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationHeader,output);
        printf("\n");
    } else {
        printf("Decoding failed\n");
    }

    uint8_t *msg_buf = NULL;
    ssize_t buf_size = encode_hdr(output, msg_buf);
    if (buf_size >= 0) {
        printf("Encoding successful\n");
    } else {
        printf("Encoding failed\n");
    }

    res = mutate_hdr(&output);
    if (res == 0) {
        printf("Mutation successful\n");
        E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationHeader,output);
        printf("\n");
    } else {
        printf("Mutation failed\n");
    }

    buf_size = encode_hdr(output, msg_buf);
    if (buf_size >= 0) {
        printf("Encoding successful\n");
    } else {
        printf("Encoding failed\n");
    }

    return 0;
}


int gen_msg(E2SM_KPM_IndicationMessage_t** output) {
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationMessage;
    void** struct_ptr = (void**)output;
    int ret  = asn_random_fill(type_descriptor, struct_ptr, 10000);
    E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationMessage,*output);
    return ret;
}

int mutate_msg(E2SM_KPM_IndicationMessage_t** output) {
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationMessage;
    void** struct_ptr = (void**)output;
    return asn_mutate(type_descriptor, struct_ptr);
}

ssize_t encode_msg(E2SM_KPM_IndicationMessage_t* input, void* buffer) {
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationMessage;
    void* struct_ptr = (void*)input;
    return aper_encode_to_new_buffer(type_descriptor, 0, struct_ptr, &buffer);
}

int decode_msg_print(void* buffer, size_t size) {
    E2SM_KPM_IndicationMessage_t* output = NULL;
    asn_dec_rval_t eval = decode_msg(buffer, size, &output);
    if (eval.code == RC_OK) {
        printf("Decoding successful\n");
        E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationMessage,output);
        free(output);
        return 0;
    } else {
        return -1;
    }
}

asn_dec_rval_t decode_msg(void* buffer, size_t size, E2SM_KPM_IndicationMessage_t** output) {
    // Prepare for decoding
    const asn_codec_ctx_t* opt_codec_ctx = NULL;  
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationMessage;  
    void** struct_ptr = (void**)output;
    const void* data_buffer = (void*)buffer;
    int skip_bits = 0;
    int unused_bits = 0;

    // Call aper_decode
    asn_dec_rval_t rval = aper_decode(opt_codec_ctx, type_descriptor, struct_ptr, data_buffer, size, 0, 0);
    return rval;
}

asn_dec_rval_t decode_msg_file(const char* filename, E2SM_KPM_IndicationMessage_t** output) {
    // Open the file
    FILE *file = fopen(filename, "rb");
    if (!file) {
        fprintf(stderr, "Failed to open file: %s\n", filename);
        exit(1);
    }

    // Get the size of the file
    fseek(file, 0, SEEK_END);
    long size = ftell(file);
    fseek(file, 0, SEEK_SET);

    // Read the file contents into a buffer
    char* buffer = (char*)malloc(size);
    if (!fread(buffer, size, 1, file)) {
        fprintf(stderr, "Failed to read file: %s\n", filename);
        exit(1);
    }

    fclose(file);

    return decode_msg(buffer, size, output);
}

int mutateMessageTest() {
    E2SM_KPM_IndicationMessage_t* output = NULL;
    int res;
    // res = gen_msg(&output);
    // if (res == 0) {
    //     std::cout << "Random gen successful\n";
    //     E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_E2SM_KPM_IndicationMessage,output);
    //     std::cout << "\n";
    // } else {
    //     std::cerr << "Random gen failed\n";
    // }

    asn_dec_rval_t eval = decode_msg_file("/home/tianchang/Desktop/asn1_test/bin_out/msg_1696603935_728564901.dat", &output);
    if (eval.code == RC_OK) {
        printf("Decoding successful\n");
        E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationMessage,output);
        printf("\n");
    } else {
        printf("Decoding failed\n");
    }

    uint8_t *msg_buf = NULL;
    ssize_t buf_size = encode_msg(output, msg_buf);
    if (buf_size >= 0) {
        printf("Encoding successful\n");
    } else {
        printf("Encoding failed\n");
    }

    res = mutate_msg(&output);
    if (res == 0) {
        printf("Mutation successful\n");
        E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_IndicationMessage,output);
        printf("\n");
    } else {
        printf("Mutation failed\n");
    }

    buf_size = encode_msg(output, msg_buf);
    if (buf_size >= 0) {
        printf("Encoding successful\n");
    } else {
        printf("Encoding failed\n");
    }

    return 0;
}


int mutateMessage(void *data, long size, void** buffer) {
    printf("Mutating message\n");
    // Prepare for decoding
    const asn_codec_ctx_t* opt_codec_ctx = NULL;  
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationMessage;
    E2SM_KPM_IndicationMessage_t* output = NULL;
    void** struct_ptr = (void**)&output;
    // const void* data_buffer = (void*)buffer;
    int skip_bits = 0;
    int unused_bits = 0;

    // Call aper_decode
    asn_dec_rval_t rval = aper_decode(opt_codec_ctx, type_descriptor, struct_ptr, data, size, 0, 0);
    if (rval.code != RC_OK) {
        return -1;
    }

    printf("Decoding successful\n");
    // E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_E2SM_KPM_IndicationMessage,output);
    // Mutate
    int res = asn_mutate(type_descriptor, struct_ptr);
    if (res != 0) {
        free(output);
        return -1;
    }

    printf("Mutation successful\n");
    // Encode
    ssize_t buf_size = aper_encode_to_new_buffer(type_descriptor, 0, *struct_ptr, buffer);
    if (buf_size < 0) {
        free(output);
        return -1;
    }
    free(output);
    printf("Encoding successful\n");
    // printBuffer(*buffer, buf_size);
    return (int)buf_size;
}

void printBuffer(void* buffer, size_t size) {
    unsigned char* byteBuffer = (unsigned char*) buffer;
    for (size_t i = 0; i < size; ++i) {
        printf("%02x ", byteBuffer[i]);  // Print each byte in hexadecimal
        if ((i + 1) % 16 == 0) {
            printf("\n");  // Print a newline every 16 bytes
        }
    }
    printf("\n");  // Print a final newline
}

int mutateHeader(void *data, long size, void* buffer) {
    printf("Mutating header\n");
    // Prepare for decoding
    const asn_codec_ctx_t* opt_codec_ctx = NULL;  
    const asn_TYPE_descriptor_t* type_descriptor = &asn_DEF_E2SM_KPM_IndicationHeader;
    E2SM_KPM_IndicationHeader_t* output = NULL;
    void** struct_ptr = (void**)&output;
    const void* data_buffer = (void*)buffer;
    int skip_bits = 0;
    int unused_bits = 0;

    // Call aper_decode
    asn_dec_rval_t rval = aper_decode(opt_codec_ctx, type_descriptor, struct_ptr, data_buffer, size, 0, 0);
    if (rval.code != RC_OK) {
        return -1;
    }

    printf("Decoding successful\n");
    // E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_E2SM_KPM_IndicationMessage,output);
    // Mutate
    int res = asn_mutate(type_descriptor, struct_ptr);
    if (res != 0) {
        return -1;
    }

    printf("Mutation successful\n");
    // E2SM_XER_PRINT(NULL,&asn_DEF_E2SM_KPM_E2SM_KPM_IndicationMessage,output);

    void *msg_buf = NULL;
    ssize_t test_size = aper_encode_to_new_buffer(type_descriptor, 0, *struct_ptr, &msg_buf);
    if (test_size < 0) {
        printf("Test encoding failed\n");
        return -1;
    }
    printf("Test encoding successful\n");

    // Encode
    ssize_t buf_size = aper_encode_to_new_buffer(type_descriptor, 0, *struct_ptr, &buffer);
    if (buf_size < 0) {
        return -1;
    }
    printf("Encoding successful\n");
    printf("buf: %s\n", (char*)buffer);
    return (int)buf_size;
}

void setRandSeed() {
    srand(time(NULL));
}

size_t get_sizeof_void_ptr() {
    return sizeof(void*);
}