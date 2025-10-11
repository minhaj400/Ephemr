package utils

import (
	"log"
	"net/smtp"

	"github.com/Minhajxdd/Ephemr/internal/config"
)

// AuthEmailUtils defines the methods for handling email-related utility functions for authentication.
type AuthEmailUtils interface {
	// buildConfirmEmailBody generates the html body of the confirmation email using a magic link.
	buildConfirmEmailBody(magicLink string) string

	// SentConfirmEmail sends a confirmation email to the specified address.
	SentConfirmEmail(to, magicLink string) error
}

type emailUtils struct{}

func NewAuthEmailUtils() AuthEmailUtils {
	return &emailUtils{}
}

func (e *emailUtils) SentConfirmEmail(to, magicLink string) error {

	from := config.Cfg.GmailId
	pass := config.Cfg.GmailAppPass

	msg := "From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: [Ephemr] Confirm Email\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		e.buildConfirmEmailBody(magicLink)

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}

func (e *emailUtils) buildConfirmEmailBody(magicLink string) string {
	return `
		<html>
		<head>
		<style>
			/* Reset / sensible defaults for many clients */
			body,table,td{margin:0;padding:0;border:0;font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif}
			img{border:0;display:block}
			a{color:inherit;text-decoration:none}
			.button{display:inline-block;padding:14px 22px;border-radius:6px;font-weight:600}
			@media only screen and (max-width:480px){
			.container{width:100% !important;padding:16px !important}
			.content{padding:18px !important}
			}
		</style>
		</head>
		<body style="background-color:#f4f6f8;">
		<!-- Preheader: hidden but visible in inbox preview -->
		<span style="color:transparent;display:none;height:0;width:0;opacity:0;visibility:hidden;overflow:hidden;">Use this link to sign in — it expires in 15 minutes.</span>

		<table role="presentation" width="100%" cellpadding="0" cellspacing="0">
			<tr>
			<td align="center" style="padding:32px 16px;">
				<table role="presentation" class="container" width="600" cellpadding="0" cellspacing="0" style="width:600px;max-width:600px;background:#ffffff;border-radius:8px;overflow:hidden;">

				<!-- Header / logo -->
				<tr>
					<td style="padding:24px 28px;border-bottom:1px solid #eef1f5;">
					<table role="presentation" width="100%">
						<tr>
						<td style="font-size:18px;font-weight:700;color:#0f1724;">
							Ephemr
						</td>
						<td align="right" style="font-size:12px;color:#6b7280;">Need help? <a href="mailto:support@example.com" style="color:#2563eb;">support@ephemr.dev</a></td>
						</tr>
					</table>
					</td>
				</tr>

				<!-- Body -->
				<tr>
					<td class="content" style="padding:32px 28px 24px;">
					<h1 style="margin:0 0 12px;font-size:22px;line-height:28px;color:#0f1724;font-weight:700;">Confirm Your Email</h1>
					<p style="margin:0 0 20px;color:#374151;line-height:1.5;font-size:15px;">We received a request to sign in to <strong>Ephemr</strong>. Click the button below to confirm sign in. This link will expire in <strong>15 minutes</strong>.</p>

					<!-- Button (use VML for Outlook fallback) -->
					<table role="presentation" cellpadding="0" cellspacing="0" style="margin:22px 0 18px;">
						<tr>
						<td align="left">
							<!--[if mso]>
							<v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="{{MAGIC_LINK}}" style="height:44px;v-text-anchor:middle;width:220px;" arcsize="8%" strokecolor="#2563eb" fillcolor="#2563eb">
							<w:anchorlock/>
							<center style="color:#ffffff;font-family:Arial, sans-serif;font-size:15px;font-weight:bold;">Sign in</center>
							</v:roundrect>
							<![endif]-->
							<!--[if !mso]><!-- -->
							<a href="` + magicLink + `" class="button" style="background:#2563eb;color:#ffffff;padding:12px 24px;border-radius:6px;display:inline-block;font-size:15px;">Sign in</a>
							<!--<![endif]-->
						</td>
						</tr>
					</table>

					<p style="margin:0 0 14px;color:#6b7280;font-size:13px;">If the button doesn't work, copy and paste this URL into your browser:</p>
					<p style="margin:0 0 18px;word-break:break-all;font-size:13px;color:#0f1724;">
						<a href=" ` + magicLink + ` " style="color:#2563eb;">` + magicLink + `</a>
					</p>

					<hr style="border:none;border-top:1px solid #eef1f5;margin:20px 0;">

					<p style="margin:0;color:#6b7280;font-size:13px;line-height:1.4;">If you didn't request this email, you can safely ignore it — no changes were made to your account.</p>
					</td>
				</tr>

				<!-- Footer -->
				<tr>
					<td style="padding:18px 28px 28px;background:#fbfdff;border-top:1px solid #eef1f5;color:#9ca3af;font-size:12px;">
					<table role="presentation" width="100%">
						<tr style="text-align:center">
						<td>© <span id="year">2025</span> Ephemr</td>
						</tr>
					</table>
					</td>
				</tr>

				</table>
			</td>
			</tr>
		</table>
	</body>
	</html>`
}
