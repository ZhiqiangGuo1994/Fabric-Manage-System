package FabricTest.saveInfo;

import java.sql.Connection;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
import java.util.UUID;

import FabricTest.user.UserContext;

public class userOperate {
	
	public void add(UserContext userContext) throws SQLException {
		/*获取数据库链接*/
		dbconfig config = new dbconfig();
		Connection con = config.getConnection();
		if(!con.isClosed()) {
			System.out.println("Create connection failure!");
		}
		Statement statement = con.createStatement();
		UUID uuid = UUID.randomUUID();
		String id = uuid.toString().replace("-", "");
		String sql = "insert into user (id,username,org,secret,mspId,delflag) values ('"+id+"','"+userContext.getName()+"','"+userContext.getAffiliation()+"','"+userContext.getSecret()+"','"+userContext.getMspId()+"','0');";
		boolean rs = statement.execute(sql);
		if(rs == false) {
			System.out.println("用户信息插入成功");
		}else{
			System.out.println("用户信息插入失败");
		}
		con.close();
	}

	public void delete(UserContext userContext) {

		
	}

	public void alert(UserContext userContext) {

	}
}