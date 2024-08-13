package main

import (
	cryptoRand "crypto/rand"
	"encoding/asn1"
	"fmt"
	"log"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"

	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-mho-go"
	e2sm_ies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go/v2/e2sm-v2-ies"
	"google.golang.org/protobuf/proto"
)

var mutatePossibility = 5

var interfaceTypes = []reflect.Type{
	reflect.TypeOf(&e2sm_mho.MhoEventTriggerDefinitionFormats_EventDefinitionFormat1{}),
	reflect.TypeOf(&e2sm_mho.E2SmMhoIndicationHeader_IndicationHeaderFormat1{}),
	reflect.TypeOf(&e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat1{}),
	reflect.TypeOf(&e2sm_mho.E2SmMhoIndicationMessage_IndicationMessageFormat2{}),
	reflect.TypeOf(&e2sm_mho.E2SmMhoControlHeader_ControlHeaderFormat1{}),
	reflect.TypeOf(&e2sm_mho.E2SmMhoControlMessage_ControlMessageFormat1{}),

	reflect.TypeOf(&e2sm_ies.Cgi_NRCgi{}),
	reflect.TypeOf(&e2sm_ies.Cgi_EUtraCgi{}),
	reflect.TypeOf(&e2sm_ies.CoreCpid_FiveGc{}),
	reflect.TypeOf(&e2sm_ies.CoreCpid_EPc{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_NG{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_XN{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_F1{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_E1{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_S1{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_X2{}),
	reflect.TypeOf(&e2sm_ies.InterfaceIdentifier_W1{}),
	reflect.TypeOf(&e2sm_ies.NodeType_GlobalEnbId{}),
	reflect.TypeOf(&e2sm_ies.NodeType_GlobalEnGnbId{}),
	reflect.TypeOf(&e2sm_ies.GroupId_FiveGc{}),
	reflect.TypeOf(&e2sm_ies.GroupId_EPc{}),
	reflect.TypeOf(&e2sm_ies.QoSid_FiveGc{}),
	reflect.TypeOf(&e2sm_ies.QoSid_EPc{}),
	reflect.TypeOf(&e2sm_ies.RrcType_Lte{}),
	reflect.TypeOf(&e2sm_ies.RrcType_Nr{}),
	reflect.TypeOf(&e2sm_ies.ServingCellArfcn_NR{}),
	reflect.TypeOf(&e2sm_ies.ServingCellArfcn_EUtra{}),
	reflect.TypeOf(&e2sm_ies.ServingCellPci_NR{}),
	reflect.TypeOf(&e2sm_ies.ServingCellPci_EUtra{}),
	reflect.TypeOf(&e2sm_ies.Ueid_GNbUeid{}),
	reflect.TypeOf(&e2sm_ies.Ueid_GNbDuUeid{}),
	reflect.TypeOf(&e2sm_ies.Ueid_GNbCuUpUeid{}),
	reflect.TypeOf(&e2sm_ies.Ueid_NgENbUeid{}),
	reflect.TypeOf(&e2sm_ies.Ueid_NgENbDuUeid{}),
	reflect.TypeOf(&e2sm_ies.Ueid_EnGNbUeid{}),
	reflect.TypeOf(&e2sm_ies.Ueid_ENbUeid{}),
	reflect.TypeOf(&e2sm_ies.EnbId_MacroENbId{}),
	reflect.TypeOf(&e2sm_ies.EnbId_HomeENbId{}),
	reflect.TypeOf(&e2sm_ies.EnbId_ShortMacroENbId{}),
	reflect.TypeOf(&e2sm_ies.EnbId_LongMacroENbId{}),
	reflect.TypeOf(&e2sm_ies.EnGnbId_EnGNbId{}),
	reflect.TypeOf(&e2sm_ies.GlobalRannodeId_GlobalGnbId{}),
	reflect.TypeOf(&e2sm_ies.GlobalRannodeId_GlobalNgEnbId{}),
	reflect.TypeOf(&e2sm_ies.GnbId_GNbId{}),
	reflect.TypeOf(&e2sm_ies.NgEnbId_MacroNgEnbId{}),
	reflect.TypeOf(&e2sm_ies.NgEnbId_ShortMacroNgEnbId{}),
	reflect.TypeOf(&e2sm_ies.NgEnbId_LongMacroNgEnbId{}),
}

func (m *Mutator) mutateStructWrapper(data []byte) ([]byte, error) {
	orig := e2sm_mho.E2SmMhoIndicationMessage{}
	err := proto.Unmarshal(data, &orig)
	if err != nil {
		return make([]byte, 0), fmt.Errorf("failed to unmarshal: %v", err)
	}
	res, err := m.mutateStruct(data, &e2sm_mho.E2SmMhoIndicationMessage{})
	if err != nil {
		mutated := e2sm_mho.E2SmMhoIndicationMessage{}
		newErr := proto.Unmarshal(res, &mutated)
		if newErr == nil {
			log.Println(fmt.Sprintf("mutation succeeded: orig: %v, mutated: %v\n", orig.String(), mutated.String()))
		}
	}
	return res, err
}

func (m *Mutator) mutateStruct(data []byte, pbMessage proto.Message) ([]byte, error) {
	err := proto.Unmarshal(data, pbMessage)
	if err != nil {
		return make([]byte, 0), fmt.Errorf("failed to unmarshal: %v", err)
	}

	mutated, err := m.chooseMutation(reflect.ValueOf(pbMessage))
	if err != nil {
		return make([]byte, 0), fmt.Errorf("failed to mutate: %v", err)
	}

	if pm, ok := mutated.Interface().(proto.Message); ok {
		return proto.Marshal(pm)
	}
	return make([]byte, 0), fmt.Errorf("mutated value does not implement proto.Message")
}

func parseTag(tag string) map[string]string {
	tagMap := make(map[string]string)
	for _, t := range strings.Split(tag, ",") {
		if strings.Contains(t, ":") {
			kv := strings.Split(t, ":")
			tagMap[kv[0]] = kv[1]
		} else {
			tagMap[t] = ""
		}
	}
	return tagMap
}

func (m *Mutator) genVal(t reflect.Type, lb, ub int, opt bool) reflect.Value {
	switch t.Kind() {
	case reflect.Interface:
		v := reflect.New(t)
		concreteTypes := make([]reflect.Type, 0)
		for _, concreteType := range interfaceTypes {
			if concreteType.Implements(t) {
				concreteTypes = append(concreteTypes, concreteType)
			}
		}
		if len(concreteTypes) == 0 {
			panic(fmt.Sprintf("No concrete types registered for interface %v", t))
		}
		concreteType := concreteTypes[m.r.Intn(len(concreteTypes))]
		v.Elem().Set(m.genVal(concreteType, lb, ub, opt))
		return v
	case reflect.Ptr:
		if opt && m.r.Intn(101) < 10 {
			return reflect.Zero(t)
		}
		v := reflect.New(t.Elem())
		v.Elem().Set(m.genVal(t.Elem(), lb, ub, opt))
		return v
	case reflect.Struct:
		if t == reflect.TypeOf(asn1.BitString{}) {
			l := m.r.Intn(1001)
			if lb != 0 || ub != 0 {
				l = m.r.Intn(ub-lb+1) + lb
			}
			b := genRandBits(l)
			v := reflect.ValueOf(asn1.BitString{Bytes: b, BitLength: l})
			return v
		}

		v := reflect.New(t)
		for i := 0; i < t.NumField(); i++ {
			// Check if the field is exported
			if t.Field(i).PkgPath != "" {
				continue
			}

			op := false
			lb, ub := 0, 0
			pbTagString := t.Field(i).Tag.Get("protobuf")
			if pbTagString != "" {
				// pbTag := parseTag(pbTagString)
				// if _, ok := pbTag["opt"]; ok {
				// 	op = true
				// }
			}

			aperTagString := t.Field(i).Tag.Get("aper")
			if aperTagString != "" {
				aperTag := parseTag(aperTagString)
				fmt.Printf("orig: %v, field: %v, string: %v, type: %v, aper tag: %v\n", t.String(), t.Field(i).Name, t.Field(i).Type.String(), t.Field(i).Type, aperTag)
				if aperTag["sizeLB"] != "" {
					sizeLB, err := strconv.Atoi(aperTag["sizeLB"])
					if err != nil {
						panic(err)
					}
					lb = sizeLB
				}
				if aperTag["sizeUB"] != "" {
					sizeUB, err := strconv.Atoi(aperTag["sizeUB"])
					if err != nil {
						panic(err)
					}
					ub = sizeUB
				}
				if aperTag["valueLB"] != "" {
					valueLB, err := strconv.Atoi(aperTag["valueLB"])
					if err != nil {
						panic(err)
					}
					lb = valueLB
				}
				if aperTag["valueUB"] != "" {
					valueUB, err := strconv.Atoi(aperTag["valueUB"])
					if err != nil {
						panic(err)
					}
					ub = valueUB
				}

				if _, ok := aperTag["optional"]; ok {
					op = true
				}
			}
			v.Field(i).Set(m.genVal(t.Field(i).Type, lb, ub, op))
		}
		return v
	case reflect.Slice:
		panic("Slice is not supported")
	case reflect.Int8:
		v := reflect.ValueOf(m.genRandInt8(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingInt8[m.r.Intn(len(interestingInt8))])
				return v
			}
			validVals := make([]int8, 0)
			for _, val := range interestingInt8 {
				if val >= int8(lb) && val <= int8(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v
	case reflect.Int16:
		v := reflect.ValueOf(m.genRandInt16(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingInt16[m.r.Intn(len(interestingInt16))])
				return v
			}
			validVals := make([]int16, 0)
			for _, val := range interestingInt16 {
				if val >= int16(lb) && val <= int16(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Int32:
		v := reflect.ValueOf(m.genRandInt32(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingInt32[m.r.Intn(len(interestingInt32))])
				return v
			}
			validVals := make([]int32, 0)
			for _, val := range interestingInt32 {
				if val >= int32(lb) && val <= int32(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Int64:
		v := reflect.ValueOf(m.genRandInt64(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingInt64[m.r.Intn(len(interestingInt64))])
				return v
			}
			validVals := make([]int64, 0)
			for _, val := range interestingInt64 {
				if val >= int64(lb) && val <= int64(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Uint8:
		v := reflect.ValueOf(m.genRandUint8(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingUint8[m.r.Intn(len(interestingUint8))])
				return v
			}
			validVals := make([]uint8, 0)
			for _, val := range interestingUint8 {
				if val >= uint8(lb) && val <= uint8(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Uint16:
		v := reflect.ValueOf(m.genRandUint16(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingUint16[m.r.Intn(len(interestingUint16))])
				return v
			}
			validVals := make([]uint16, 0)
			for _, val := range interestingUint16 {
				if val >= uint16(lb) && val <= uint16(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Uint32:
		v := reflect.ValueOf(m.genRandUint32(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingUint32[m.r.Intn(len(interestingUint32))])
				return v
			}
			validVals := make([]uint32, 0)
			for _, val := range interestingUint32 {
				if val >= uint32(lb) && val <= uint32(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Uint64:
		v := reflect.ValueOf(m.genRandUint64(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingUint64[m.r.Intn(len(interestingUint64))])
				return v
			}
			validVals := make([]uint64, 0)
			for _, val := range interestingUint64 {
				if val >= uint64(lb) && val <= uint64(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return v
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Float32:
		v := reflect.ValueOf(m.genRandFloat32(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingFloat32[m.r.Intn(len(interestingFloat32))])
				return v
			}
			validVals := make([]float32, 0)
			for _, val := range interestingFloat32 {
				if val >= float32(lb) && val <= float32(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return reflect.ValueOf(float32(m.rr.Float32()))
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Float64:
		v := reflect.ValueOf(m.genRandFloat64(lb, ub))
		switch m.r.Intn(2) {
		case 0:
			if lb == 0 && ub == 0 {
				v = reflect.ValueOf(interestingFloat64[m.r.Intn(len(interestingFloat64))])
				return v
			}
			validVals := make([]float64, 0)
			for _, val := range interestingFloat64 {
				if val >= float64(lb) && val <= float64(ub) {
					validVals = append(validVals, val)
				}
			}
			if len(validVals) == 0 {
				return reflect.ValueOf(float64(m.rr.Float64()))
			}
			v = reflect.ValueOf(validVals[m.r.Intn(len(validVals))])
		}
		return v

	case reflect.Complex64:
		v := reflect.ValueOf(interestingComplex64[m.r.Intn(len(interestingComplex64))])
		return v
	case reflect.Complex128:
		v := reflect.ValueOf(interestingComplex128[m.r.Intn(len(interestingComplex128))])
		return v
	case reflect.String:
		v := reflect.ValueOf(interestingStrings[m.r.Intn(len(interestingStrings))])
		return v
	case reflect.Bool:
		switch m.r.Intn(2) {
		case 0:
			return reflect.ValueOf(true)
		}
		return reflect.ValueOf(false)
	default:
		panic(fmt.Sprintf("Unsupported type: %v", t))
	}
	return reflect.Value{}
}

func (m *Mutator) chooseMutation(v reflect.Value) (reflect.Value, error) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		// unexported field
		if v.Field(i).Type().PkgPath() != "" {
			continue
		}

		op := false
		lb, ub := 0, 0
		pbTagString := v.Type().Field(i).Tag.Get("protobuf")
		if pbTagString != "" {
			// pbTag := parseTag(pbTagString)
			// if _, ok := pbTag["opt"]; ok {
			// 	op = true
			// }
		}

		aperTagString := v.Type().Field(i).Tag.Get("aper")
		if aperTagString != "" {
			aperTag := parseTag(aperTagString)
			if aperTag["sizeLB"] != "" {
				sizeLB, err := strconv.Atoi(aperTag["sizeLB"])
				if err != nil {
					panic(err)
				}
				lb = sizeLB
			}
			if aperTag["sizeUB"] != "" {
				sizeUB, err := strconv.Atoi(aperTag["sizeUB"])
				if err != nil {
					panic(err)
				}
				ub = sizeUB
			}
			if aperTag["valueLB"] != "" {
				valueLB, err := strconv.Atoi(aperTag["valueLB"])
				if err != nil {
					panic(err)
				}
				lb = valueLB
			}
			if aperTag["valueUB"] != "" {
				valueUB, err := strconv.Atoi(aperTag["valueUB"])
				if err != nil {
					panic(err)
				}
				ub = valueUB
			}

			if _, ok := aperTag["optional"]; ok {
				op = true
			}
		}

		if m.r.Intn(101) < mutatePossibility {
			fieldVal := m.filedMutate(field, lb, ub, op)
			v.Field(i).Set(fieldVal)
			continue
		}
		if field.Kind() == reflect.Struct || field.Kind() == reflect.Pointer ||
			field.Kind() == reflect.Interface {
			newV, err := m.chooseMutation(field)
			if err != nil {
				return reflect.Value{}, err
			}
			v.Field(i).Set(newV)
		}
	}
	return v, nil
}

func (m *Mutator) filedMutate(v reflect.Value, lb int, ub int, opt bool) reflect.Value {
	t := v.Type()
	switch t.Kind() {
	case reflect.Interface:
		concreteTypes := make([]reflect.Type, 0)
		for _, concreteType := range interfaceTypes {
			if concreteType.Implements(t) {
				concreteTypes = append(concreteTypes, concreteType)
			}
		}
		if len(concreteTypes) == 0 {
			panic(fmt.Sprintf("No concrete types registered for interface %v", t))
		}
		concreteType := concreteTypes[m.r.Intn(len(concreteTypes))]
		v.Set(m.filedMutate(reflect.New(concreteType).Elem(), lb, ub, opt))
		return v
	case reflect.Ptr:
		if v.IsNil() {
			ptr := reflect.New(t.Elem())
			ptr.Elem().Set(m.filedMutate(ptr.Elem(), lb, ub, opt))
			v.Set(ptr)
		} else {
			v.Set(m.filedMutate(v.Elem(), lb, ub, opt))
		}
		return v
		// if opt {
		// 	v.Set(reflect.Zero(t))
		// 	return v
		// }
	case reflect.Struct:
		if t == reflect.TypeOf(asn1.BitString{}) {
			orig := make([]byte, 0)
			origBits := 0
			for i := 0; i < v.NumField(); i++ {
				if t.Field(i).PkgPath != "" {
					continue
				}
				if t.Field(i).Type == reflect.TypeOf(make([]byte, 0)) {
					orig = v.Field(i).Bytes()
				} else if t.Field(i).Type == reflect.TypeOf(uint32(0)) {
					origBits = int(v.Field(i).Uint())
				} else {
					panic(fmt.Sprintf("BitString field %v has unknown type %v", t.Field(i).Name, t.Field(i).Type))
				}
			}
			newByte, newLen := m.mutateAsn1Bits(lb, ub, orig, origBits)

			bitStringVal := reflect.New(t).Elem()
			for i := 0; i < bitStringVal.NumField(); i++ {
				if t.Field(i).PkgPath != "" {
					continue
				}
				if t.Field(i).Type == reflect.TypeOf(make([]byte, 0)) {
					bitStringVal.Field(i).Set(reflect.ValueOf(newByte))
				} else if t.Field(i).Type == reflect.TypeOf(uint32(0)) {
					bitStringVal.Field(i).Set(reflect.ValueOf(uint32(newLen)))
				} else {
					panic(fmt.Sprintf("BitString field %v has unknown type %v", t.Field(i).Name, t.Field(i).Type))
				}
			}
			v.Set(bitStringVal)
			return v
		}

		for i := 0; i < t.NumField(); i++ {
			// Check if the field is exported
			if t.Field(i).PkgPath != "" {
				continue
			}

			op := false
			lb, ub = 0, 0
			pbTagString := t.Field(i).Tag.Get("protobuf")
			if pbTagString != "" {
				// pbTag := parseTag(pbTagString)
				// if _, ok := pbTag["opt"]; ok {
				// 	op = true
				// }
			}

			aperTagString := t.Field(i).Tag.Get("aper")
			if aperTagString != "" {
				aperTag := parseTag(aperTagString)
				fmt.Printf("orig: %v, field: %v, string: %v, type: %v, aper tag: %v\n", t.String(), t.Field(i).Name, t.Field(i).Type.String(), t.Field(i).Type, aperTag)
				if aperTag["sizeLB"] != "" {
					sizeLB, err := strconv.Atoi(aperTag["sizeLB"])
					if err != nil {
						panic(err)
					}
					lb = sizeLB
				}
				if aperTag["sizeUB"] != "" {
					sizeUB, err := strconv.Atoi(aperTag["sizeUB"])
					if err != nil {
						panic(err)
					}
					ub = sizeUB
				}
				if aperTag["valueLB"] != "" {
					valueLB, err := strconv.Atoi(aperTag["valueLB"])
					if err != nil {
						panic(err)
					}
					lb = valueLB
				}
				if aperTag["valueUB"] != "" {
					valueUB, err := strconv.Atoi(aperTag["valueUB"])
					if err != nil {
						panic(err)
					}
					ub = valueUB
				}

				if _, ok := aperTag["optional"]; ok {
					op = true
				}
			}
			fieldVal := m.filedMutate(v.Field(i), lb, ub, op)
			v.Field(i).Set(fieldVal)
		}
		return v
	case reflect.Slice:
		if v.IsNil() {
			v.Set(reflect.MakeSlice(t, 0, 0))
		}
		switch m.r.Intn(4) {
		case 0: // generating new array
			l := m.r.Intn(101)
			if lb != 0 || ub != 0 {
				l = m.r.Intn(ub-lb+1) + lb
			}
			arr := reflect.MakeSlice(t, l, l)
			for i := 0; i < l; i++ {
				arr.Index(i).Set(m.genVal(t.Elem(), 0, 0, opt))
			}
			v.Set(arr)
			return v
		case 1: // change val in array
			if v.Len() == 0 {
				// append
				v.Set(reflect.Append(v, m.genVal(t.Elem(), 0, 0, opt)))
			}
			idx := m.r.Intn(v.Len())
			v.Index(idx).Set(m.genVal(t.Elem(), 0, 0, opt))
			return v
		case 2: // delete val in array
			if v.Len() == 0 {
				return v
			}
			idx := m.r.Intn(v.Len())
			v.Set(reflect.AppendSlice(v.Slice(0, idx), v.Slice(idx+1, v.Len())))
			return v
		case 3: // insert val in array
			if v.Len() == 0 {
				v.Set(reflect.Append(v, m.genVal(t.Elem(), 0, 0, opt)))
				return v
			}
			idx := m.r.Intn(v.Len())
			v.Set(reflect.AppendSlice(v.Slice(0, idx), reflect.Append(v.Slice(idx, v.Len()), m.genVal(t.Elem(), 0, 0, opt))))
			return v
		}

	case reflect.Int8:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingInt8[m.r.Intn(len(interestingInt8))]
			if val == int8(v.Int()) || ((lb != 0 || ub != 0) && (val < int8(lb) || val > int8(ub))) {
				val = m.genRandInt8(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandInt8(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Int()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += int64(m.r.Intn(10) + 1)
			} else {
				val -= int64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v
	case reflect.Int16:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingInt16[m.r.Intn(len(interestingInt16))]
			if val == int16(v.Int()) || ((lb != 0 || ub != 0) && (val < int16(lb) || val > int16(ub))) {
				val = m.genRandInt16(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandInt16(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Int()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += int64(m.r.Intn(10) + 1)
			} else {
				val -= int64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v
	case reflect.Int32:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingInt32[m.r.Intn(len(interestingInt32))]
			if val == int32(v.Int()) || ((lb != 0 || ub != 0) && (val < int32(lb) || val > int32(ub))) {
				val = m.genRandInt32(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandInt32(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Int()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += int64(m.r.Intn(10) + 1)
			} else {
				val -= int64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v

	case reflect.Int64:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingInt64[m.r.Intn(len(interestingInt64))]
			if val == v.Int() || ((lb != 0 || ub != 0) && (val < int64(lb) || val > int64(ub))) {
				val = m.genRandInt64(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandInt64(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Int()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += int64(m.r.Intn(10) + 1)
			} else {
				val -= int64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v

	case reflect.Uint8:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingUint8[m.r.Intn(len(interestingUint8))]
			if val == uint8(v.Uint()) || ((lb != 0 || ub != 0) && (val < uint8(lb) || val > uint8(ub))) {
				val = m.genRandUint8(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandUint8(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Uint()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += uint64(m.r.Intn(10) + 1)
			} else {
				val -= uint64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v

	case reflect.Uint16:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingUint16[m.r.Intn(len(interestingUint16))]
			if val == uint16(v.Uint()) || ((lb != 0 || ub != 0) && (val < uint16(lb) || val > uint16(ub))) {
				val = m.genRandUint16(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandUint16(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Uint()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += uint64(m.r.Intn(10) + 1)
			} else {
				val -= uint64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v

	case reflect.Uint32:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingUint32[m.r.Intn(len(interestingUint32))]
			if val == uint32(v.Uint()) || ((lb != 0 || ub != 0) && (val < uint32(lb) || val > uint32(ub))) {
				val = m.genRandUint32(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandUint32(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Uint()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += uint64(m.r.Intn(10) + 1)
			} else {
				val -= uint64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v
	case reflect.Uint64:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingUint64[m.r.Intn(len(interestingUint64))]
			if val == v.Uint() || ((lb != 0 || ub != 0) && (val < uint64(lb) || val > uint64(ub))) {
				val = m.genRandUint64(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandUint64(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Uint()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += uint64(m.r.Intn(10) + 1)
			} else {
				val -= uint64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v

	case reflect.Float32:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingFloat32[m.r.Intn(len(interestingFloat32))]
			if val == float32(v.Float()) || ((lb != 0 || ub != 0) && (val < float32(lb) || val > float32(ub))) {
				val = m.genRandFloat32(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandFloat32(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Float()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += float64(m.r.Intn(10) + 1)
			} else {
				val -= float64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v

	case reflect.Float64:
		r := m.r.Intn(3)
		switch r {
		case 0:
			val := interestingFloat64[m.r.Intn(len(interestingFloat64))]
			if val == v.Float() || ((lb != 0 || ub != 0) && (val < float64(lb) || val > float64(ub))) {
				val = m.genRandFloat64(lb, ub)
			}
			v.Set(reflect.ValueOf(val))
		case 1:
			val := m.genRandFloat64(lb, ub)
			v.Set(reflect.ValueOf(val))
		case 2:
			val := v.Float()
			mode := m.r.Intn(2) // 0: add, 1: sub
			if mode == 0 {
				val += float64(m.r.Intn(10) + 1)
			} else {
				val -= float64(m.r.Intn(10) + 1)
			}
			v.Set(reflect.ValueOf(val))
		}
		return v
	case reflect.Complex64:
		val := interestingComplex64[m.r.Intn(len(interestingComplex64))]
		v.Set(reflect.ValueOf(val))
		return v
	case reflect.Complex128:
		val := interestingComplex128[m.r.Intn(len(interestingComplex128))]
		v.Set(reflect.ValueOf(val))
		return v
	case reflect.String:
		str := m.mutateString(v.String())
		v.Set(reflect.ValueOf(str))
		return v
	case reflect.Bool:
		b := v.Bool()
		v.Set(reflect.ValueOf(!b))
		return v
	default:
		panic(fmt.Sprintf("Unsupported type: %v", t))
	}
	return v
}

func (m *Mutator) mutateString(orig string) string {
	r := m.r.Intn(6)
	switch r {
	case 0:
		pos := m.r.Intn(len(orig))
		replacement := interestingStrings[m.r.Intn(len(interestingStrings))]
		return orig[:pos] + replacement + orig[pos+1:]
	case 1:
		pos := m.r.Intn(len(orig))
		insertValue := interestingStrings[m.r.Intn(len(interestingStrings))]
		return orig[:pos] + insertValue + orig[pos:]
	case 2:
		if len(orig) == 0 {
			return orig
		}
		pos := m.r.Intn(len(orig))
		return orig[:pos] + orig[pos+1:]
	case 3:
		runeSlice := []rune(orig)
		rand.Shuffle(len(runeSlice), func(i, j int) {
			runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
		})
		return string(runeSlice)
	case 4:
		addition := interestingStrings[m.r.Intn(len(interestingStrings))]
		if m.r.Intn(2) == 0 {
			return orig + addition
		}
		return addition + orig
	case 5:
		return orig + m.generateRandomString(m.r.Intn(10)+1)
	}
	return orig
}

func (m *Mutator) generateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = byte(m.r.Intn(256))
	}
	return string(result)
}

func (m *Mutator) mutateAsn1Bits(lb, ub int, orig []byte, origBits int) ([]byte, int) {
	if lb > ub {
		panic("Lower bound cannot be greater than upper bound")
	}

	r := m.r.Intn(3)
	switch r {
	case 0: // random generation
		numBits := lb
		if ub != lb {
			numBits = lb + m.r.Intn(ub-lb+1)
		}

		data := genRandBits(numBits)
		return data, numBits

	case 1: // bits flip
		fullBytes := origBits / 8
		remainingBits := origBits % 8
		for i := 0; i < fullBytes; i++ {
			orig[i] = ^orig[i]
		}

		if remainingBits > 0 {
			orig[fullBytes] = ^(orig[fullBytes] & ^(1<<remainingBits - 1))
		}
		return orig, origBits

	case 2: // add/remove bits
		if lb == ub {
			return m.mutateAsn1Bits(lb, ub, orig, origBits) // try other mutations
		}
		newLen := m.r.Intn(ub-lb+1) + lb
		diff := newLen - origBits
		if diff < 0 {
			bytesToKeep := newLen / 8
			bitsToKeepInLastByte := newLen % 8

			if bitsToKeepInLastByte > 0 {
				// Zero out the trailing bits
				mask := byte(0xFF) << (8 - bitsToKeepInLastByte)
				orig[bytesToKeep] &= mask
				bytesToKeep++
			}
			orig = orig[:bytesToKeep]
		} else {
			if origBits%8 != 0 {
				for i := origBits % 8; i < 8; i++ {
					if diff == 0 {
						break
					}
					if m.r.Intn(2) == 1 {
						orig[len(orig)-1] |= (1 << (8 - i - 1))
					}
					diff--

				}
			}
			extra := genRandBits(diff)
			orig = append(orig, extra...)
		}
		return orig, newLen
	}
	return orig, origBits
}

func genRandBits(len int) []byte {
	numBytes := int(math.Ceil(float64(len) / 8))
	data := make([]byte, numBytes)
	cryptoRand.Read(data)

	// Ensure trailing bits in the last byte are zero
	unusedBits := 8 - len%8
	if unusedBits == 8 {
		unusedBits = 0
	}
	unusedMask := (1 << unusedBits) - 1
	data[numBytes-1] &= ^byte(unusedMask)
	return data
}

var (
	// For integers (signed):
	interestingInt8  = []int8{-128, -1, 0, 1, 16, 32, 64, 100, 127}
	interestingInt16 = []int16{-32768, -128, -1, 0, 1, 16, 32, 64, 127, 255, 256, 32767}
	interestingInt32 = []int32{-2147483648, -32768, -128, -1, 0, 1, 16, 32, 64, 127, 255, 256, 32767, 2147483647}
	interestingInt64 = []int64{-9223372036854775808, -2147483648, -32768, -128, -1, 0, 1, 16, 32, 64, 127, 255, 256, 32767, 2147483647, 9223372036854775807}

	// For integers (unsigned):
	interestingUint8  = []uint8{0, 1, 16, 32, 64, 100, 127, 255}
	interestingUint16 = []uint16{0, 1, 16, 32, 64, 127, 255, 256, 32767, 65535}
	interestingUint32 = []uint32{0, 1, 16, 32, 64, 127, 255, 256, 32767, 65535, 2147483647, 4294967295}
	interestingUint64 = []uint64{0, 1, 16, 32, 64, 127, 255, 256, 32767, 65535, 2147483647, 4294967295, 9223372036854775807, 18446744073709551615}

	// For floats:
	interestingFloat32 = []float32{-math.MaxFloat32, -math.SmallestNonzeroFloat32, -1.0, 0.0, math.SmallestNonzeroFloat32, 1.0, math.MaxFloat32}
	interestingFloat64 = []float64{-math.MaxFloat64, -math.SmallestNonzeroFloat64, -1.0, 0.0, math.SmallestNonzeroFloat64, 1.0, math.MaxFloat64}

	// For complex:
	interestingComplex64  = []complex64{0, 1 + 1i, -1 - 1i, math.MaxFloat32, -math.MaxFloat32, 1i, -1i}
	interestingComplex128 = []complex128{0, 1 + 1i, -1 - 1i, math.MaxFloat64, -math.MaxFloat64, 1i, -1i}

	// For strings (considering edge cases):
	interestingStrings = []string{"", " ", "\n", "\t", "\000", "a", "A", "1", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "[", "]", "{", "}", "|", "\\", ":", ";", "'", "\"", ",", "<", ".", ">", "/", "?"}

	// For bool:
	interestingBool = []bool{true, false}
)
