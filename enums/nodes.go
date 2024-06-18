package enums

type Node string

func CreateNode(grpcNode string) Node {
	return Node(grpcNode)
}

const (
	MAIN_NODE           Node = "grpc.trongrid.io:50051"
	SHASTA_NODE         Node = "grpc.shasta.trongrid.io:50051"
	NILE_NODE           Node = "grpc.nile.trongrid.io:50051"
	TRONEX_NODE         Node = "grpc.nile.trongrid.io:50051"
	CUSTOM_TEST_1_NODE  Node = "161.117.83.38:50051"
	CUSTOM_TEST_2_NODE  Node = "3.225.171.164:50051"
	CUSTOM_TEST_3_NODE  Node = "52.53.189.99:50051"
	CUSTOM_TEST_4_NODE  Node = "18.196.99.16:50051"
	CUSTOM_TEST_5_NODE  Node = "34.253.187.192:50051"
	CUSTOM_TEST_6_NODE  Node = "18.133.82.227:50051"
	CUSTOM_TEST_7_NODE  Node = "35.180.51.163:50051"
	CUSTOM_TEST_8_NODE  Node = "54.252.224.209:50051"
	CUSTOM_TEST_9_NODE  Node = "18.228.15.36:50051"
	CUSTOM_TEST_10_NODE Node = "52.15.93.92:50051"
	CUSTOM_TEST_11_NODE Node = "34.220.77.106:50051"
	CUSTOM_TEST_12_NODE Node = "15.207.144.3:50051"
	CUSTOM_TEST_13_NODE Node = "13.124.62.58:50051"
	CUSTOM_TEST_14_NODE Node = "15.222.19.181:50051"
	CUSTOM_TEST_15_NODE Node = "18.209.42.127:50051"
	CUSTOM_TEST_16_NODE Node = "3.218.137.187:50051"
	CUSTOM_TEST_17_NODE Node = "34.237.210.82:50051"
	CUSTOM_TEST_18_NODE Node = "47.241.20.47:50051"
	CUSTOM_TEST_19_NODE Node = "161.117.85.97:50051"
	CUSTOM_TEST_20_NODE Node = "161.117.224.116:50051"
)