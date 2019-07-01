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
import java.util.Collection;
import java.util.Iterator;
import java.util.logging.Level;
import java.util.logging.Logger;

import FabricTest.client.FabricClient;
import FabricTest.config.FabricConfig;
import FabricTest.user.UserContext;
import FabricTest.util.Util;
import org.hyperledger.fabric.sdk.Channel;
import org.hyperledger.fabric.sdk.ChannelConfiguration;
import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric.sdk.Orderer;
import org.hyperledger.fabric.sdk.Peer;
import org.hyperledger.fabric.sdk.security.CryptoSuite;

/**
 * 
 * @author Balaji Kadambi
 *
 */

public class CreateChannel {

	public static void main(String[] args) {
		try {
			CryptoSuite.Factory.getCryptoSuite();
			Util.cleanUp();
			// Construct Channel
			UserContext org1Admin = new UserContext();
			File pkFolder1 = new File(FabricConfig.ORG1_USR_ADMIN_PK);
			File[] pkFiles1 = pkFolder1.listFiles();
			File certFolder1 = new File(FabricConfig.ORG1_USR_ADMIN_CERT);
			File[] certFiles1 = certFolder1.listFiles();
			Enrollment enrollOrg1Admin = Util.getEnrollment(FabricConfig.ORG1_USR_ADMIN_PK, pkFiles1[0].getName(),
					FabricConfig.ORG1_USR_ADMIN_CERT, certFiles1[0].getName());
			org1Admin.setEnrollment(enrollOrg1Admin);
			org1Admin.setMspId(FabricConfig.ORG1_MSP);
			org1Admin.setName(FabricConfig.ADMIN);

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

			// Create a new channel
			Orderer orderer = fabClient.getInstance().newOrderer(FabricConfig.ORDERER_NAME, FabricConfig.ORDERER_URL);
			ChannelConfiguration channelConfiguration = new ChannelConfiguration(new File(FabricConfig.CHANNEL_CONFIG_PATH));

			byte[] channelConfigurationSignatures = fabClient.getInstance()
					.getChannelConfigurationSignature(channelConfiguration, org1Admin);

			Channel mychannel = fabClient.getInstance().newChannel(FabricConfig.CHANNEL_NAME, orderer, channelConfiguration,
					channelConfigurationSignatures);

			Peer peer0_org1 = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_0, FabricConfig.ORG1_PEER_0_URL);
			Peer peer1_org1 = fabClient.getInstance().newPeer(FabricConfig.ORG1_PEER_1, FabricConfig.ORG1_PEER_1_URL);
			Peer peer0_org2 = fabClient.getInstance().newPeer(FabricConfig.ORG2_PEER_0, FabricConfig.ORG2_PEER_0_URL);
			Peer peer1_org2 = fabClient.getInstance().newPeer(FabricConfig.ORG2_PEER_1, FabricConfig.ORG2_PEER_1_URL);
			
			Peer peer0_org3 = fabClient.getInstance().newPeer(FabricConfig.ORG3_PEER_0, FabricConfig.ORG3_PEER_0_URL);
			Peer peer1_org3 = fabClient.getInstance().newPeer(FabricConfig.ORG3_PEER_1, FabricConfig.ORG3_PEER_1_URL);
			Peer peer0_org4 = fabClient.getInstance().newPeer(FabricConfig.ORG4_PEER_0, FabricConfig.ORG4_PEER_0_URL);
			Peer peer1_org4 = fabClient.getInstance().newPeer(FabricConfig.ORG4_PEER_1, FabricConfig.ORG4_PEER_1_URL);

			mychannel.joinPeer(peer0_org1);
			mychannel.joinPeer(peer1_org1);
			
			mychannel.addOrderer(orderer);

			mychannel.initialize();
			
			fabClient.getInstance().setUserContext(org2Admin);
			mychannel = fabClient.getInstance().getChannel(FabricConfig.CHANNEL_NAME);
			mychannel.joinPeer(peer0_org2);
			mychannel.joinPeer(peer1_org2);
			
			fabClient.getInstance().setUserContext(org3Admin);
			mychannel = fabClient.getInstance().getChannel(FabricConfig.CHANNEL_NAME);
			mychannel.joinPeer(peer0_org3);
			mychannel.joinPeer(peer1_org3);
			
			fabClient.getInstance().setUserContext(org4Admin);
			mychannel = fabClient.getInstance().getChannel(FabricConfig.CHANNEL_NAME);
			mychannel.joinPeer(peer0_org4);
			mychannel.joinPeer(peer1_org4);
			
			Logger.getLogger(CreateChannel.class.getName()).log(Level.INFO, "Channel created "+mychannel.getName());
            Collection peers = mychannel.getPeers();
            Iterator peerIter = peers.iterator();
            while (peerIter.hasNext())
            {
            	  Peer pr = (Peer) peerIter.next();
            	  Logger.getLogger(CreateChannel.class.getName()).log(Level.INFO,pr.getName()+ " at " + pr.getUrl());
            }
            
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}