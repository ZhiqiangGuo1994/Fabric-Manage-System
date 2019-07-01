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

import java.util.Collection;
import java.util.HashMap;
import java.util.Map;

import FabricTest.client.CAClient;
import FabricTest.client.ChannelClient;
import FabricTest.client.FabricClient;
import FabricTest.config.FabricConfig;
import FabricTest.user.UserContext;
import FabricTest.util.Util;
import org.hyperledger.fabric.sdk.ChaincodeID;
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

public class InvokeQueryChaincode {

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
			Peer peer = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_0, FabricConfig.ORG1_PEER_0_URL);
			EventHub eventHub = fabClient.getInstance().newEventHub("eventhub01", "grpc://localhost:7053");
			Orderer orderer = fabClient.getInstance().newOrderer(FabricConfig.ORDERER_NAME, FabricConfig.ORDERER_URL);
			channel.addPeer(peer);
			channel.addEventHub(eventHub);
			channel.addOrderer(orderer);
			channel.initialize();
			
			long begintime = System.currentTimeMillis();
			for(int i=0;i<5000;i++) {
				String temp = "专利" + Integer.toString(i);
				String[] args1 = {temp};
				Collection<ProposalResponse>  responses1Query = null;
				try {
					responses1Query = channelClient.queryByChainCode("patent_chaincode", "readPatent", args1);
				}catch(Exception e) {
					continue;
				}
				for (ProposalResponse pres : responses1Query) {
					String stringResponse = null;
					try {
						stringResponse = new String(pres.getChaincodeActionResponsePayload());
					}catch(Exception e) {
						continue;
					}
					System.out.println(stringResponse);
				}
			}
			long endtime = System.currentTimeMillis();
			System.out.println("消耗时间：" + (endtime-begintime));
			
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
