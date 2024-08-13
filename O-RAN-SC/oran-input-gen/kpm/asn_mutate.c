/*
 * Tianchang Yang <tzy5088@psu.edu>
 */
#include <asn_internal.h>
#include <asn_mutate.h>
#include <constr_TYPE.h>

int
asn_mutate(const struct asn_TYPE_descriptor_s *td, void **struct_ptr) {
    if(td && td->op->mutate) {
        asn_mutate_result_t res =
            td->op->mutate(td, struct_ptr, 0);
        return (res.code == MUTATE_OK) ? 0 : -1;
    } else {
        return -1;
    }
}
