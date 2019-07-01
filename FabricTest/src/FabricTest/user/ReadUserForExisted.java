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

import FabricTest.client.CAClient;
import FabricTest.config.FabricConfig;
import FabricTest.util.Util;

/**
 * 
 * @author Balaji Kadambi
 *
 */

public class ReadUserForExisted {

	public static void main(String args[]) {
		try {
			String caUrl = FabricConfig.CA_ORG1_URL;
			CAClient caClient = new CAClient(caUrl, null);
			UserContext adminuserContext = Util.readUserContext("org1", "admin");
			caClient.setAdminUserContext(adminuserContext);
			
			System.out.println();
			// Register and Enroll user to Org3MSP

			UserContext userContext1 = new UserContext();
			String name1 = "user" + System.currentTimeMillis();
			userContext1.setName(name1);
			userContext1.setAffiliation(FabricConfig.ORG1);
			userContext1.setMspId(FabricConfig.ORG1_MSP);

			String eSecret1 = caClient.registerUser(name1, FabricConfig.ORG1);
			userContext1 = caClient.enrollUser(userContext1, eSecret1);
			
			userContext1 = caClient.reenrollUser(userContext1, eSecret1);


		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
