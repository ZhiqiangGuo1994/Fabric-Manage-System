/****************************************************** 
 *  Copyright 2018 IBM Corporation 
 *  Licensed under the Apache License, Version 2.0 (the "License"); 
 *  you may not use this file except in compliance with the License. 
 *  You may obtain a copy of the License at 
 *  http://www.apache.org/licenses/LICENSE-2.0 
 *  Unless required by applicable law or agreed to in writing, software 
 *  distributed under the License is distributed on an "AS IS" BASIS, 
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. 
 *  See the License for the specific language governing permissions and 
 *  limitations under the License.
 */
package FabricTest.chaincode.invocation;

import static java.nio.charset.StandardCharsets.UTF_8;

import java.io.File;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

import FabricTest.client.CAClient;
import FabricTest.client.ChannelClient;
import FabricTest.client.FabricClient;
import FabricTest.config.FabricConfig;
import FabricTest.user.UserContext;
import FabricTest.util.Util;
import org.hyperledger.fabric.sdk.ChaincodeID;
import org.hyperledger.fabric.sdk.ChaincodeResponse.Status;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.EventHub;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.ProposalResponse;
import org.hyperledger.fabric.sdk.TransactionProposalRequest;

/**
 * 
 * @author Balaji Kadambi
 *
 */

public class InvokeChaincode {

	private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
	private static final String EXPECTED_EVENT_NAME = "event";

	public static void main(String args[]) {
		try {
			Util.cleanUp();
			String caUrl = FabricConfig.CA_ORG1_URL;
			CAClient caClient = new CAClient(caUrl, null);
			// Enroll Admin to Org1MSP
			UserContext adminUserContext = new UserContext();
			adminUserContext.setName(FabricConfig.ADMIN);
			adminUserContext.setAffiliation(FabricConfig.ORG1);
			adminUserContext.setMspId(FabricConfig.ORG1_MSP);
			caClient.setAdminUserContext(adminUserContext);
			adminUserContext = caClient.enrollAdminUser(FabricConfig.ADMIN, FabricConfig.ADMIN_PASSWORD);

			FabricClient fabClient = new FabricClient(adminUserContext);

			ChannelClient channelClient = fabClient.createChannelClient(FabricConfig.CHANNEL_NAME);
			Channel channel = channelClient.getChannel();
			Peer peer1 = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_0, FabricConfig.ORG1_PEER_0_URL);
			EventHub eventHub1 = fabClient.getInstance().newEventHub("eventhub01", "grpc://localhost:7053");
			
			/*Peer peer2 = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_1, FabricConfig.ORG1_PEER_1_URL);
			EventHub eventHub2 = fabClient.getInstance().newEventHub("eventhub01", "grpc://localhost:7058");*/
			Orderer orderer = fabClient.getInstance().newOrderer(FabricConfig.ORDERER_NAME, FabricConfig.ORDERER_URL);
			channel.addPeer(peer1);
			channel.addEventHub(eventHub1);
			
			/*channel.addPeer(peer2);
			channel.addEventHub(eventHub2);*/
			channel.addOrderer(orderer);
			channel.initialize();

			long begintime = System.currentTimeMillis();
			int flag = 0; 
			for(int i=0;i<5000;i++) {
				String x = "专利" + Integer.toString(i);
				
				System.out.println(x);
				TransactionProposalRequest request = fabClient.getInstance().newTransactionProposalRequest();
				ChaincodeID ccid = ChaincodeID.newBuilder().setName(FabricConfig.CHAINCODE_1_NAME).build();
				request.setChaincodeID(ccid);
				request.setFcn("initPatent");
				String[] arguments = {x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x};
				/*String[] arguments = {"c785f57f7e6b44678fb35151bd2423a7","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1"};*/
				/*String[] arguments = {"专利2","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1","专利1"};*/
				/*String[] arguments = {"87862fcd9149414abda4e3370e879360"};*/
				request.setArgs(arguments);
				request.setProposalWaitTime(1000);

				Map<String, byte[]> tm2 = new HashMap<>();
				tm2.put("HyperLedgerFabric", "TransactionProposalRequest:JavaSDK".getBytes(UTF_8));
				tm2.put("method", "TransactionProposalRequest".getBytes(UTF_8));
				tm2.put("result", ":)".getBytes(UTF_8));
				tm2.put(EXPECTED_EVENT_NAME, EXPECTED_EVENT_DATA);
				request.setTransientMap(tm2);
				Collection<ProposalResponse> responses = null;
				try {
					responses = channelClient.sendTransactionProposal(request);
				}catch(Exception e) {
					flag++;
					continue;
				}
				for (ProposalResponse res : responses) {
					Status status = res.getStatus();
					if(status.name().equals("SUCCESS")) {
						Logger.getLogger(InvokeChaincode.class.getName()).log(Level.INFO,
								"Invoked initPatent on " + FabricConfig.CHAINCODE_1_NAME + ". Status - " + status);
					}else {
						flag++;
						continue;
					}
				}
				
				/*String[] args2= {};
				Collection<ProposalResponse>  responses1Query = channelClient.queryByChainCode("patent_chaincode", "getCreator", args2);
				for (ProposalResponse pres : responses1Query) {
					String stringResponse = new String(pres.getChaincodeActionResponsePayload());
					System.out.println(stringResponse);
				}*/
			}
			long endtime = System.currentTimeMillis();
			System.out.println("消耗时间：" + (endtime-begintime));
			System.out.println("提案失败次数：" + flag);
			
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
