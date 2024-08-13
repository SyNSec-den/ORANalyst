/*
 * Tianchang Yang <tzy5088@psu.edu>
 */
#ifndef	ASN_MUTATE
#define	ASN_MUTATE

/* Forward declarations */
struct asn_TYPE_descriptor_s;
struct asn_encoding_constraints_s;

typedef struct asn_mutate_result_s {
    enum {
        MUTATE_FAILED = -1, 
        MUTATE_OK = 0,
    } code;
} asn_mutate_result_t;
typedef asn_mutate_result_t (asn_mutate_f)(
    const struct asn_TYPE_descriptor_s *td, void **struct_ptr,
    const struct asn_encoding_constraints_s *memb_constraints);

/*
 * Returns 0 if the structure was properly initialized, -1 otherwise.
 */
int asn_mutate(const struct asn_TYPE_descriptor_s *td, void **struct_ptr);

#endif	/* ASN_MUTATE */
