package FabricTest.saveInfo;

import java.sql.*;

public class dbconfig {

	Connection con;
	
	public Connection getConnection() {
		try {
			Class.forName("com.mysql.jdbc.Driver");
			System.out.println("数据库驱动加载成功");
		}catch(ClassNotFoundException e) {
			e.printStackTrace();
		}
		
		try {
			con = DriverManager.getConnection("jdbc:mysql:"+"//120.77.144.63:3306/blockchain?serverTimezone=GMT%2B8","ubuntu","guozhiqiang");
			System.out.println("数据库连接成功");
		}catch(SQLException e) {
			e.printStackTrace();
		}
		
		return con;
	}

	public static void main(String[] args) {
		new dbconfig().getConnection();
	}
}
