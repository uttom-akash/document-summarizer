using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;

namespace auth.Controllers;

[Route("api/[controller]")]
[ApiController]
public class AuthController : ControllerBase
{
    private readonly ILogger<AuthController> _logger;

    public AuthController(ILogger<AuthController> logger)
    {
        _logger = logger;
    }

    [HttpGet("login")]
    public  string Login(string username, string role)
    {
        // throw new Exception("my error!");
        
        var key = "MySuperSecretKey1234518918918991891i91819819191910910910910101!";

        var credentials = new SigningCredentials(
            new SymmetricSecurityKey(Encoding.UTF8.GetBytes(key)),
            SecurityAlgorithms.HmacSha256);
        
        var claims = new List<Claim>
        {
            new Claim("UserName", username),
            new Claim(ClaimTypes.Role, role)
        };
        
        var token = new JwtSecurityToken(
            issuer: "authapi",
            audience: "user",
            claims: claims,
            expires: DateTime.UtcNow.AddMinutes(5),
            signingCredentials: credentials);

        var jwttoken = new JwtSecurityTokenHandler().WriteToken(token);
        
        return jwttoken;
    }
    
    [Authorize(policy:"RequireAdminRole")]
    [HttpGet("logout")]
    public ActionResult Logout()
    {
        _logger.LogInformation("logout");
        return Ok();
    }
}