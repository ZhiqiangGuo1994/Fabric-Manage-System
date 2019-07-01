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
package FabricTest.user;

import java.io.File;

import org.hyperledger.fabric.sdk.Enrollment;
import org.hyperledger.fabric_ca.sdk.helper.Config;

import FabricTest.client.CAClient;
import FabricTest.config.FabricConfig;
import FabricTest.saveInfo.userOperate;
import FabricTest.util.Util;

/**
 * 
 * @author Balaji Kadambi
 *
 */

public class RegisterEnrollUser {

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

			String caUrl2 = FabricConfig.CA_ORG2_URL;
			CAClient caClient2 = new CAClient(caUrl2, null);
			/* Enroll Admin to Org2MSP */
			UserContext adminUserContext2 = new UserContext();
			adminUserContext2.setName(FabricConfig.ADMIN);
			adminUserContext2.setAffiliation(FabricConfig.ORG2);
			adminUserContext2.setMspId(FabricConfig.ORG2_MSP);
			caClient2.setAdminUserContext(adminUserContext2);
			adminUserContext2 = caClient2.enrollAdminUser(FabricConfig.ADMIN, FabricConfig.ADMIN_PASSWORD);

			String caUrl3 = FabricConfig.CA_ORG3_URL;
			CAClient caClient3 = new CAClient(caUrl3, null);
			/* Enroll Admin to Org2MSP */
			UserContext adminUserContext3 = new UserContext();
			adminUserContext3.setName(FabricConfig.ADMIN);
			adminUserContext3.setAffiliation(FabricConfig.ORG3);
			adminUserContext3.setMspId(FabricConfig.ORG3_MSP);
			caClient3.setAdminUserContext(adminUserContext3);
			adminUserContext3 = caClient3.enrollAdminUser(FabricConfig.ADMIN, FabricConfig.ADMIN_PASSWORD);

			String caUrl4 = FabricConfig.CA_ORG4_URL;
			CAClient caClient4 = new CAClient(caUrl4, null);
			/* Enroll Admin to Org4MSP */
			UserContext adminUserContext4 = new UserContext();
			adminUserContext4.setName(FabricConfig.ADMIN);
			adminUserContext4.setAffiliation(FabricConfig.ORG4);
			adminUserContext4.setMspId(FabricConfig.ORG4_MSP);
			caClient4.setAdminUserContext(adminUserContext4);
			adminUserContext4 = caClient4.enrollAdminUser(FabricConfig.ADMIN, FabricConfig.ADMIN_PASSWORD);
			
			System.out.println();
			// Register and Enroll user to Org3MSP

			UserContext userContext3 = new UserContext();
			String name3 = "user" + System.currentTimeMillis();
			userContext3.setName(name3);
			userContext3.setAffiliation(FabricConfig.ORG3);
			userContext3.setMspId(FabricConfig.ORG3_MSP);

			String eSecret3 = caClient3.registerUser(name3, FabricConfig.ORG3);
			userContext3 = caClient3.enrollUser(userContext3, eSecret3);

			System.out.println();

			// Register and Enroll user to Org4MSP
			UserContext userContext4 = new UserContext();
			String name4 = "user" + System.currentTimeMillis();
			userContext4.setName(name4);
			userContext4.setAffiliation(FabricConfig.ORG4);
			userContext4.setMspId(FabricConfig.ORG4_MSP);

			String eSecret4 = caClient4.registerUser(name4, FabricConfig.ORG4);
			userContext4 = caClient4.enrollUser(userContext4, eSecret4);

		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
