package utils

func GetUserAccessEmailTemplate(machineId string, userName string, password string) string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Vithsutra Technologies Smart Telephone System</title>
</head>
<body style="margin: 0; padding: 0; font-family: 'Poppins', Arial, sans-serif; background-color: #f4f4f4;">
    <table role="presentation" style="width: 100%; border-collapse: collapse;">
        <tr>
            <td style="padding: 0;">
                <table role="presentation" style="width: 100%; max-width: 600px; margin: 0 auto; background-color: #ffffff; border-radius: 8px; margin-top: 20px; margin-bottom: 20px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
                    <!-- Modern Header Design -->
                    <tr>
                        <td style="background: linear-gradient(135deg, #4169E1 0%, #5c88ff 100%); padding: 40px 30px; border-radius: 8px 8px 0 0; text-align: center;">
                            <div style="margin-bottom: 20px;">
                                <!-- Wave Design -->
                                <div style="margin-bottom: 15px;">
                                    <svg width="100" height="24" viewBox="0 0 100 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                        <path d="M0 12C16.6667 12 16.6667 24 33.3333 24C50 24 50 12 66.6667 12C83.3333 12 83.3333 24 100 24V0H0V12Z" fill="rgba(255,255,255,0.1)"/>
                                    </svg>
                                </div>
                                <h1 style="color: #ffffff; font-size: 28px; margin: 0; font-weight: 600; font-family: 'Poppins', Arial, sans-serif; text-transform: uppercase; letter-spacing: 1px;">
                                    Vithsutra Technologies
                                </h1>
                            </div>
                            <!-- Welcome Message with Modern Design -->
                            <div style="background: rgba(255,255,255,0.1); border-radius: 12px; padding: 25px; margin-top: 20px;">
                                <h2 style="color: #ffffff; font-size: 24px; margin: 0 0 15px 0; font-family: 'Poppins', Arial, sans-serif;">
                                    Welcome to Smart Telephone System
                                </h2>
                                <div style="background: rgba(255,255,255,0.15); border-radius: 8px; padding: 12px; margin: 0 auto; max-width: 300px;">
                                    <p style="color: #ffffff; font-size: 16px; margin: 0 0 5px 0; font-family: 'Poppins', Arial, sans-serif;">
                                        POS Machine ID
                                    </p>
                                    <p style="color: #ffffff; font-size: 24px; font-weight: 600; margin: 0; font-family: 'Poppins', Arial, sans-serif; letter-spacing: 2px;">
                                        ` + machineId + `
                                    </p>
                                </div>
                            </div>
                        </td>
                    </tr>

                    <!-- Content -->
                    <tr>
                        <td style="padding: 40px 30px;">
                            <!-- Welcome Info -->
                            <div style="background: #f8f9fa; border-radius: 12px; padding: 25px; margin-bottom: 30px; border-left: 4px solid #4169E1;">
                                <p style="color: #666666; font-size: 16px; line-height: 24px; margin: 0; font-family: 'Poppins', Arial, sans-serif;">
                                    Below are your login credentials for the assigned POS machine to manage RFID card recharges. These credentials are specifically bound to your assigned POS machine.
                                </p>
                            </div>
                            
                            <!-- Credentials Box with Modern Design -->
                            <div style="background: linear-gradient(to right, #ffffff, #f8f9fa); border-radius: 12px; padding: 25px; margin: 20px 0; box-shadow: 0 2px 8px rgba(65,105,225,0.1); border: 1px solid rgba(65,105,225,0.2);">
                                <h3 style="color: #4169E1; font-size: 18px; margin: 0 0 20px 0; font-family: 'Poppins', Arial, sans-serif; display: flex; align-items: center;">
                                    <span style="background: #4169E1; color: white; padding: 6px; border-radius: 6px; font-size: 14px; margin-right: 10px;">🔐</span>
                                    Login Credentials
                                </h3>
                                <table role="presentation" style="width: 100%; border-collapse: collapse;">
                                    <tr>
                                        <td style="padding: 12px 15px; background: rgba(65,105,225,0.05); border-radius: 6px;">
                                            <p style="margin: 0; color: #666666; font-family: 'Poppins', Arial, sans-serif;">Username</p>
                                            <p style="margin: 5px 0 0 0; color: #333333; font-weight: 600; font-family: 'Poppins', Arial, sans-serif; font-size: 16px;">` + userName + `</p>
                                        </td>
                                    </tr>
                                    <tr>
                                        <td style="padding: 12px 15px; margin-top: 10px; display: block; background: rgba(65,105,225,0.05); border-radius: 6px;">
                                            <p style="margin: 0; color: #666666; font-family: 'Poppins', Arial, sans-serif;">Password</p>
                                            <p style="margin: 5px 0 0 0; color: #333333; font-weight: 600; font-family: 'Poppins', Arial, sans-serif; font-size: 16px;">` + password + `</p>
                                        </td>
                                    </tr>
                                </table>
                                <p style="color: #ff6b6b; font-size: 14px; margin: 15px 0 0 0; font-family: 'Poppins', Arial, sans-serif; display: flex; align-items: center;">
                                    <span style="color: #ff6b6b; margin-right: 5px;">⚠️</span>
                                    Please change your password upon first login for security
                                </p>
                            </div>

                            <!-- Important Notice -->
                            <div style="background: #fff4e5; border-radius: 12px; padding: 20px; margin-top: 30px; border-left: 4px solid #ff9800;">
                                <h4 style="color: #ff9800; font-size: 16px; margin: 0 0 10px 0; font-family: 'Poppins', Arial, sans-serif;">
                                    Important Notice
                                </h4>
                                <p style="color: #666666; font-size: 14px; margin: 0; font-family: 'Poppins', Arial, sans-serif; line-height: 1.6;">
                                    These credentials will only work with POS Machine ID: ` + machineId + `. Any attempt to use these credentials on a different machine will not be successful.
                                </p>
                            </div>
                        </td>
                    </tr>

                    <!-- Modern Footer -->
                    <tr>
                        <td style="padding: 30px; background: linear-gradient(135deg, #f8f9fa 0%, #ffffff 100%); border-radius: 0 0 8px 8px; text-align: center;">
                            <p style="color: #666666; font-size: 14px; margin: 0 0 10px 0; font-family: 'Poppins', Arial, sans-serif;">
                                Need help? Our support team is available 24/7
                            </p>
                            <p style="color: #999999; font-size: 14px; margin: 0; font-family: 'Poppins', Arial, sans-serif;">
                                © 2024 Vithsutra Technologies | Smart Telephone Solutions
                            </p>
                        </td>
                    </tr>
                </table>
            </td>
        </tr>
    </table>
</body>
</html>`
}
