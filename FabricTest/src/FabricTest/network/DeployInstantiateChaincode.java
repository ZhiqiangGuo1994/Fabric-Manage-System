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
package FabricTest.network;

import java.io.File;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.logging.Level;
import java.util.logging.Logger;

import FabricTest.client.ChannelClient;
import FabricTest.client.FabricClient;
import FabricTest.config.FabricConfig;
import FabricTest.user.UserContext;
import FabricTest.util.Util;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.ProposalResponse;
import org.hyperledger.fabric.sdk.TransactionRequest.Type;
import org.hyperledger.fabric.sdk.security.CryptoSuite;

/**
 * 
 * @author Balaji Kadambi
 *
 */

public class DeployInstantiateChaincode {

	public static void main(String[] args) {
		try {
			CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();

			UserContext org1Admin = new UserContext();
			File pkFolder1 = new File(FabricConfig.ORG1_USR_ADMIN_PK);
			File[] pkFiles1 = pkFolder1.listFiles();
			File certFolder = new File(FabricConfig.ORG1_USR_ADMIN_CERT);
			File[] certFiles = certFolder.listFiles();
			Enrollment enrollOrg1Admin = Util.getEnrollment(FabricConfig.ORG1_USR_ADMIN_PK, pkFiles1[0].getName(),
					FabricConfig.ORG1_USR_ADMIN_CERT, certFiles[0].getName());
			org1Admin.setEnrollment(enrollOrg1Admin);
			org1Admin.setMspId("Org1MSP");
			org1Admin.setName("admin");

			UserContext org2Admin = new UserContext();
			File pkFolder2 = new File(FabricConfig.ORG2_USR_ADMIN_PK);
			File[] pkFiles2 = pkFolder2.listFiles();
			File certFolder2 = new File(FabricConfig.ORG2_USR_ADMIN_CERT);
			File[] certFiles2 = certFolder2.listFiles();
			Enrollment enrollOrg2Admin = Util.getEnrollment(FabricConfig.ORG2_USR_ADMIN_PK, pkFiles2[0].getName(),
					FabricConfig.ORG2_USR_ADMIN_CERT, certFiles2[0].getName());
			org2Admin.setEnrollment(enrollOrg2Admin);
			org2Admin.setMspId(FabricConfig.ORG2_MSP);
			org2Admin.setName(FabricConfig.ADMIN);

			UserContext org3Admin = new UserContext();
			File pkFolder3 = new File(FabricConfig.ORG3_USR_ADMIN_PK);
			File[] pkFiles3 = pkFolder3.listFiles();
			File certFolder3 = new File(FabricConfig.ORG3_USR_ADMIN_CERT);
			File[] certFiles3 = certFolder3.listFiles();
			Enrollment enrollOrg3Admin = Util.getEnrollment(FabricConfig.ORG3_USR_ADMIN_PK, pkFiles3[0].getName(),
					FabricConfig.ORG3_USR_ADMIN_CERT, certFiles3[0].getName());
			org3Admin.setEnrollment(enrollOrg3Admin);
			org3Admin.setMspId(FabricConfig.ORG3_MSP);
			org3Admin.setName(FabricConfig.ADMIN);

			UserContext org4Admin = new UserContext();
			File pkFolder4 = new File(FabricConfig.ORG4_USR_ADMIN_PK);
			File[] pkFiles4 = pkFolder4.listFiles();
			File certFolder4 = new File(FabricConfig.ORG4_USR_ADMIN_CERT);
			File[] certFiles4 = certFolder4.listFiles();
			Enrollment enrollOrg4Admin = Util.getEnrollment(FabricConfig.ORG4_USR_ADMIN_PK, pkFiles4[0].getName(),
					FabricConfig.ORG4_USR_ADMIN_CERT, certFiles4[0].getName());
			org4Admin.setEnrollment(enrollOrg4Admin);
			org4Admin.setMspId(FabricConfig.ORG4_MSP);
			org4Admin.setName(FabricConfig.ADMIN);

			FabricClient fabClient = new FabricClient(org1Admin);

			Channel mychannel = fabClient.getInstance().newChannel(FabricConfig.CHANNEL_NAME);
			Orderer orderer = fabClient.getInstance().newOrderer(FabricConfig.ORDERER_NAME, FabricConfig.ORDERER_URL);
			Peer peer0_org1 = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_0, FabricConfig.ORG1_PEER_0_URL);
			Peer peer1_org1 = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_1, FabricConfig.ORG1_PEER_1_URL);
			Peer peer0_org2 = fabClient.getInstance().newPeer(FabricConfig.ORG2_PEER_0, FabricConfig.ORG2_PEER_0_URL);
			Peer peer1_org2 = fabClient.getInstance().newPeer(FabricConfig.ORG2_PEER_1, FabricConfig.ORG2_PEER_1_URL);
			Peer peer0_org3 = fabClient.getInstance().newPeer(FabricConfig.ORG3_PEER_0, FabricConfig.ORG3_PEER_0_URL);
			Peer peer1_org3 = fabClient.getInstance().newPeer(FabricConfig.ORG3_PEER_1, FabricConfig.ORG3_PEER_1_URL);
			Peer peer0_org4 = fabClient.getInstance().newPeer(FabricConfig.ORG4_PEER_0, FabricConfig.ORG4_PEER_0_URL);
			Peer peer1_org4 = fabClient.getInstance().newPeer(FabricConfig.ORG4_PEER_1, FabricConfig.ORG4_PEER_1_URL);
			mychannel.addOrderer(orderer);
			mychannel.addPeer(peer0_org1);
			mychannel.addPeer(peer1_org1);
			mychannel.addPeer(peer0_org2);
			mychannel.addPeer(peer1_org2);
			mychannel.addPeer(peer0_org3);
			mychannel.addPeer(peer1_org3);
			mychannel.addPeer(peer0_org4);
			mychannel.addPeer(peer1_org4);
			mychannel.initialize();

			List<Peer> org1Peers = new ArrayList<Peer>();
			org1Peers.add(peer0_org1);
			org1Peers.add(peer1_org1);

			List<Peer> org2Peers = new ArrayList<Peer>();
			org2Peers.add(peer0_org2);
			org2Peers.add(peer1_org2);

			List<Peer> org3Peers = new ArrayList<Peer>();
			org3Peers.add(peer0_org3);
			org3Peers.add(peer1_org3);

			List<Peer> org4Peers = new ArrayList<Peer>();
			org4Peers.add(peer0_org4);
			org4Peers.add(peer1_org4);

			Collection<ProposalResponse> response = fabClient.deployChainCode(FabricConfig.CHAINCODE_1_NAME,
					FabricConfig.CHAINCODE_1_PATH, FabricConfig.CHAINCODE_ROOT_DIR, Type.GO_LANG.toString(),
					FabricConfig.CHAINCODE_1_VERSION, org1Peers);

			for (ProposalResponse res : response) {
				Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
						FabricConfig.CHAINCODE_1_NAME + "- Chain code deployment " + res.getStatus());
			}

			fabClient.getInstance().setUserContext(org2Admin);

			response = fabClient.deployChainCode(FabricConfig.CHAINCODE_1_NAME, FabricConfig.CHAINCODE_1_PATH,
					FabricConfig.CHAINCODE_ROOT_DIR, Type.GO_LANG.toString(), FabricConfig.CHAINCODE_1_VERSION,
					org2Peers);

			for (ProposalResponse res : response) {
				Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
						FabricConfig.CHAINCODE_1_NAME + "- Chain code deployment " + res.getStatus());
			}

			fabClient.getInstance().setUserContext(org3Admin);

			response = fabClient.deployChainCode(FabricConfig.CHAINCODE_1_NAME, FabricConfig.CHAINCODE_1_PATH,
					FabricConfig.CHAINCODE_ROOT_DIR, Type.GO_LANG.toString(), FabricConfig.CHAINCODE_1_VERSION,
					org3Peers);

			for (ProposalResponse res : response) {
				Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
						FabricConfig.CHAINCODE_1_NAME + "- Chain code deployment " + res.getStatus());
			}
			
			fabClient.getInstance().setUserContext(org4Admin);

			response = fabClient.deployChainCode(FabricConfig.CHAINCODE_1_NAME, FabricConfig.CHAINCODE_1_PATH,
					FabricConfig.CHAINCODE_ROOT_DIR, Type.GO_LANG.toString(), FabricConfig.CHAINCODE_1_VERSION,
					org4Peers);

			for (ProposalResponse res : response) {
				Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
						FabricConfig.CHAINCODE_1_NAME + "- Chain code deployment " + res.getStatus());
			}

			ChannelClient channelClient = new ChannelClient(mychannel.getName(), mychannel, fabClient);

			String[] arguments = { "" };
			
			response = channelClient.instantiateChainCode(FabricConfig.CHAINCODE_1_NAME,
					FabricConfig.CHAINCODE_1_VERSION, FabricConfig.CHAINCODE_1_PATH, Type.GO_LANG.toString(), "init",
					arguments,null);
			
			/*add chaincode endorsement policies*/
			/*response = channelClient.instantiateChainCode(FabricConfig.CHAINCODE_1_NAME,
					FabricConfig.CHAINCODE_1_VERSION, FabricConfig.CHAINCODE_1_PATH, Type.GO_LANG.toString(), "init",
					arguments, (FabricConfig.ENDORSEMENTPOLICY_ROOT_DIR +File.separator + FabricConfig.ENDORSEMENTPOLICY_1_NAME));*/

			for (ProposalResponse res : response) {
				Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
						FabricConfig.CHAINCODE_1_NAME + "- Chain code instantiation " + res.getStatus());
			}

		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
