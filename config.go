package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fatih/color"
	"gopkg.in/ini.v1"
)

var (
	config = defaultConfiguration()
)

//#region Credentials

var (
	placeholderToken    string = "REPLACE_WITH_YOUR_TOKEN_OR_DELETE_LINE"
	placeholderEmail    string = "REPLACE_WITH_YOUR_EMAIL_OR_DELETE_LINE"
	placeholderPassword string = "REPLACE_WITH_YOUR_PASSWORD_OR_DELETE_LINE"
)

type configurationCredentials struct {
	// Login
	Token    string `json:"token,omitempty"`    // required for bot token (this or login)
	Email    string `json:"email,omitempty"`    // required for login (this or token)
	Password string `json:"password,omitempty"` // required for login (this or token)
	UserBot  bool   `json:"userBot,omitempty"`  // required
	// APIs
	TwitterAccessToken         string `json:"twitterAccessToken,omitempty"`         // optional
	TwitterAccessTokenSecret   string `json:"twitterAccessTokenSecret,omitempty"`   // optional
	TwitterConsumerKey         string `json:"twitterConsumerKey,omitempty"`         // optional
	TwitterConsumerSecret      string `json:"twitterConsumerSecret,omitempty"`      // optional
	FlickrApiKey               string `json:"flickrApiKey,omitempty"`               // optional
	GoogleDriveCredentialsJSON string `json:"googleDriveCredentialsJSON,omitempty"` // optional
}

//#endregion

//#region Configuration

// cd = Config Default
// Needed for settings used without redundant nil checks, and settings defaulting + creation
var (
	// Setup
	cdDebugOutput          bool   = false
	cdCommandPrefix        string = "ddg "
	cdAllowSkipping        bool   = true
	cdScanOwnMessages      bool   = false
	cdCheckPermissions     bool   = true
	cdAllowGlobalCommands  bool   = true
	cdGithubUpdateChecking bool   = true
	// Appearance
	cdPresenceEnabled bool               = true
	cdPresenceStatus  string             = string(discordgo.StatusIdle)
	cdPresenceType    discordgo.GameType = discordgo.GameTypeGame
	cdInflateCount    int64              = 0
)

func defaultConfiguration() configuration {
	return configuration{
		// Required
		Credentials: configurationCredentials{
			Token:    placeholderToken,
			Email:    placeholderEmail,
			Password: placeholderPassword,
		},
		// Setup
		Admins:                         []string{},
		DebugOutput:                    cdDebugOutput,
		CommandPrefix:                  cdCommandPrefix,
		AllowSkipping:                  cdAllowSkipping,
		ScanOwnMessages:                cdScanOwnMessages,
		CheckPermissions:               cdCheckPermissions,
		AllowGlobalCommands:            cdAllowGlobalCommands,
		AutorunHistory:                 false,
		AsynchronousHistory:            false,
		DownloadRetryMax:               3,
		DownloadTimeout:                60,
		GithubUpdateChecking:           cdGithubUpdateChecking,
		FilterDuplicateImages:          false,
		FilterDuplicateImagesThreshold: 0,
		// Appearance
		PresenceEnabled:      cdPresenceEnabled,
		PresenceStatus:       cdPresenceStatus,
		PresenceType:         cdPresenceType,
		FilenameDateFormat:   "2006-01-02_15-04-05 ",
		InflateCount:         &cdInflateCount,
		NumberFormatEuropean: false,
	}
}

type configuration struct {
	// Required
	Credentials configurationCredentials `json:"credentials"` // required
	// Setup
	Admins                         []string                    `json:"admins"`                                   // optional
	AdminChannels                  []configurationAdminChannel `json:"adminChannels"`                            // optional
	DebugOutput                    bool                        `json:"debugOutput"`                              // optional, defaults
	CommandPrefix                  string                      `json:"commandPrefix"`                            // optional, defaults
	AllowSkipping                  bool                        `json:"allowSkipping"`                            // optional, defaults
	ScanOwnMessages                bool                        `json:"scanOwnMessages"`                          // optional, defaults
	CheckPermissions               bool                        `json:"checkPermissions,omitempty"`               // optional, defaults
	AllowGlobalCommands            bool                        `json:"allowGlobalCommmands,omitempty"`           // optional, defaults
	AutorunHistory                 bool                        `json:"autorunHistory,omitempty"`                 // optional, defaults
	AsynchronousHistory            bool                        `json:"asyncHistory,omitempty"`                   // optional, defaults
	DownloadRetryMax               int                         `json:"downloadRetryMax,omitempty"`               // optional, defaults
	DownloadTimeout                int                         `json:"downloadTimeout,omitempty"`                // optional, defaults
	GithubUpdateChecking           bool                        `json:"githubUpdateChecking"`                     // optional, defaults
	FilterDuplicateImages          bool                        `json:"filterDuplicateImages,omitempty"`          // optional, defaults
	FilterDuplicateImagesThreshold float64                     `json:"filterDuplicateImagesThreshold,omitempty"` // optional, defaults
	// Appearance
	PresenceEnabled          bool               `json:"presenceEnabled"`                    // optional, defaults
	PresenceStatus           string             `json:"presenceStatus"`                     // optional, defaults
	PresenceType             discordgo.GameType `json:"presenceType,omitempty"`             // optional, defaults
	PresenceOverwrite        *string            `json:"presenceOverwrite,omitempty"`        // optional, unused if undefined
	PresenceOverwriteDetails *string            `json:"presenceOverwriteDetails,omitempty"` // optional, unused if undefined
	PresenceOverwriteState   *string            `json:"presenceOverwriteState,omitempty"`   // optional, unused if undefined
	FilenameDateFormat       string             `json:"filenameDateFormat,omitempty"`       // optional, defaults
	EmbedColor               *string            `json:"embedColor,omitempty"`               // optional, defaults to role if undefined, then defaults random if no role color
	InflateCount             *int64             `json:"inflateCount,omitempty"`             // optional, defaults to 0 if undefined
	NumberFormatEuropean     bool               `json:"numberFormatEuropean,omitempty"`     // optional, defaults
	// Channels
	AllChannels          *configurationChannel  `json:"allChannels,omitempty"`          // optional, defaults
	AllChannelsBlacklist *[]string              `json:"allChannelsBlacklist,omitempty"` // optional
	Channels             []configurationChannel `json:"channels"`                       // required

	/* IDEAS / TODO:

	*

	 */
}

//#endregion

//#region Channels

// ccd = Channel Config Default
// Needed for settings used without redundant nil checks, and settings defaulting + creation
var (
	// Setup
	ccdEnabled       bool = true
	ccdAllowCommands bool = true
	ccdErrorMessages bool = true
	ccdScanEdits     bool = true
	ccdIgnoreBots    bool = false
	// Appearance
	ccdUpdatePresence           bool     = true
	ccdReactWhenDownloaded      bool     = true
	ccdReactWhenDownloadedEmoji string   = ""
	ccdBlacklistReactEmojis     []string = []string{}
	ccdTypeWhileProcessing      bool     = false
	// Rules for Access
	ccdUsersAllWhitelisted bool = true
	// Rules for Saving
	ccdDivideFoldersByServer  bool = false
	ccdDivideFoldersByChannel bool = false
	ccdDivideFoldersByUser    bool = false
	ccdDivideFoldersByType    bool = true
	ccdSaveImages             bool = true
	ccdSaveVideos             bool = true
	ccdSaveAudioFiles         bool = false
	ccdSaveTextFiles          bool = false
	ccdSaveOtherFiles         bool = false
	ccdSavePossibleDuplicates bool = false
	ccdExtensionBlacklist          = []string{
		".htm",
		".html",
		".php",
		".exe",
		".dll",
		".bin",
		".cmd",
		".sh",
		".py",
		".jar",
	}
)

type configurationChannel struct {
	// Main
	ChannelID   string    `json:"channel"`            // required
	ChannelIDs  *[]string `json:"channels,omitempty"` // alternative to ChannelID
	Destination string    `json:"destination"`        // required
	// Setup
	Enabled                 *bool `json:"enabled,omitempty"`                 // optional, defaults
	AllowCommands           *bool `json:"allowCommands,omitempty"`           // optional, defaults
	ErrorMessages           *bool `json:"errorMessages,omitempty"`           // optional, defaults
	ScanEdits               *bool `json:"scanEdits,omitempty"`               // optional, defaults
	IgnoreBots              *bool `json:"ignoreBots,omitempty"`              // optional, defaults
	OverwriteAutorunHistory *bool `json:"overwriteAutorunHistory,omitempty"` // optional
	// Appearance
	UpdatePresence           *bool     `json:"updatePresence,omitempty"`           // optional, defaults
	ReactWhenDownloaded      *bool     `json:"reactWhenDownloaded,omitempty"`      // optional, defaults
	ReactWhenDownloadedEmoji *string   `json:"reactWhenDownloadedEmoji,omitempty"` // optional, defaults
	BlacklistReactEmojis     *[]string `json:"blacklistReactEmojis,omitempty"`     // optional
	TypeWhileProcessing      *bool     `json:"typeWhileProcessing,omitempty"`      // optional, defaults
	// Overwrite Global Settings
	OverwriteFilenameDateFormat *string `json:"overwriteFilenameDateFormat,omitempty"` // optional
	OverwriteAllowSkipping      *bool   `json:"overwriteAllowSkipping,omitempty"`      // optional
	OverwriteEmbedColor         *string `json:"overwriteEmbedColor,omitempty"`         // optional, defaults to role if undefined, then defaults random if no role color
	// Rules for Access
	UsersAllWhitelisted *bool     `json:"usersAllWhitelisted,omitempty"` // optional, defaults to true
	UserWhitelist       *[]string `json:"userWhitelist,omitempty"`       // optional, only relevant if above is false
	UserBlacklist       *[]string `json:"userBlacklist,omitempty"`       // optional
	// Rules for Saving
	DivideFoldersByServer  *bool     `json:"divideFoldersByServer,omitempty"`  // optional, defaults
	DivideFoldersByChannel *bool     `json:"divideFoldersByChannel,omitempty"` // optional, defaults
	DivideFoldersByUser    *bool     `json:"divideFoldersByUser,omitempty"`    // optional, defaults
	DivideFoldersByType    *bool     `json:"divideFoldersByType,omitempty"`    // optional, defaults
	SaveImages             *bool     `json:"saveImages,omitempty"`             // optional, defaults
	SaveVideos             *bool     `json:"saveVideos,omitempty"`             // optional, defaults
	SaveAudioFiles         *bool     `json:"saveAudioFiles,omitempty"`         // optional, defaults
	SaveTextFiles          *bool     `json:"saveTextFiles,omitempty"`          // optional, defaults
	SaveOtherFiles         *bool     `json:"saveOtherFiles,omitempty"`         // optional, defaults
	SavePossibleDuplicates *bool     `json:"savePossibleDuplicates,omitempty"` // optional, defaults
	ExtensionBlacklist     *[]string `json:"extensionBlacklist,omitempty"`     // optional, defaults
	DomainBlacklist        *[]string `json:"domainBlacklist,omitempty"`        // optional, defaults
	SaveAllLinksToFile     *string   `json:"saveAllLinksToFile,omitempty"`     // optional

	/* IDEAS / TODO:

	// These require an efficient way to check roles. I haven't really looked into it.
	* RolesAllWhitelisted *bool     `json:"rolesAllWhitelisted,omitempty"` // optional, defaults to true
	* RoleWhitelist       *[]string `json:"roleWhitelist,omitempty"`       // optional
	* RoleBlacklist       *[]string `json:"roleBlacklist,omitempty"`       // optional

	*/
}

//#endregion

//#region Admin Channels

type configurationAdminChannel struct {
	// Required
	ChannelID string `json:"channel"` // required

	/* IDEAS / TODO:

	* UnrestrictAdminCommands *bool `json:"unrestrictAdminCommands,omitempty"` // optional, defaults
	* SendErrorLogs *bool `json:"sendErrorLogs,omitempty"` // optional
	* SendHourlyDigest *bool `json:"sendHourlyDigest,omitempty"` // optional
	* SendDailyDigest *bool `json:"sendDailyDigest,omitempty"` // optional

	 */
}

//#endregion

func loadConfig() {
	// Load settings
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(color.HiRedString("Failed to open settings file...\t%s", err))
		createConfig()
		properExit()
	} else {
		fixed := string(configContent)
		// Fix backslashes
		fixed = strings.ReplaceAll(fixed, "\\", "\\\\")
		for strings.Contains(fixed, "\\\\\\") {
			fixed = strings.ReplaceAll(fixed, "\\\\\\", "\\\\")
		}
		//TODO: Not even sure if this is realistic to do but would be nice to have line comma & trailing comma fixing
		// Re-convert
		configContent = []byte(fixed)
		// Parse
		newConfig := defaultConfiguration()
		err = json.Unmarshal(configContent, &newConfig)
		if err != nil {
			log.Println(color.HiRedString("Failed to parse settings file...\t%s", err))
			log.Println(logPrefixHelper, color.MagentaString("Please ensure you're following proper JSON format syntax."))
			properExit()
		}
		config = newConfig

		// Channel Config Defaults
		// this is dumb but don't see a better way to initialize defaults
		for i := 0; i < len(config.Channels); i++ {
			channelDefault(&config.Channels[i])
		}
		if config.AllChannels != nil {
			channelDefault(config.AllChannels)
		}

		// Debug Output
		if config.DebugOutput {
			s, err := json.MarshalIndent(config, "", "\t")
			if err != nil {
				log.Println(logPrefixDebug, color.HiRedString("Failed to output settings...\t%s", err))
			} else {
				log.Println(logPrefixDebug, color.HiYellowString("Parsed Fixed Settings into JSON:\n\n"),
					color.YellowString(string(s)),
				)
			}
		}

		// Credentials Check
		if (config.Credentials.Token == "" || config.Credentials.Token == placeholderToken) &&
			(config.Credentials.Email == "" || config.Credentials.Email == placeholderEmail) &&
			(config.Credentials.Password == "" || config.Credentials.Password == placeholderPassword) {
			log.Println(color.HiRedString("No valid discord login found. Token, Email, and Password are all invalid..."))
			log.Println(color.HiYellowString("Please save your credentials & info into \"%s\" then restart...", configPath))
			log.Println(logPrefixHelper, color.MagentaString("If your credentials are already properly saved, please ensure you're following proper JSON format syntax."))
			log.Println(logPrefixHelper, color.MagentaString("You DO NOT NEED `Token` *AND* `Email`+`Password`, just one OR the other."))
			properExit()
		}
	}
}

func createConfig() {
	log.Println(logPrefixSetup, color.YellowString("Creating new settings file..."))

	enteredToken := placeholderToken
	enteredEmail := placeholderEmail
	enteredPassword := placeholderPassword

	enteredAdmin := "REPLACE_WITH_YOUR_DISCORD_USER_ID"

	enteredBaseChannel := "REPLACE_WITH_DISCORD_CHANNEL_ID_TO_DOWNLOAD_FROM"
	enteredBaseDestination := "REPLACE_WITH_FOLDER_LOCATION_TO_DOWNLOAD_TO"

	// Separate from Defaultconfiguration because there's some elements we want to strip for settings creation
	defaultConfig := configuration{
		Credentials: configurationCredentials{
			Token:    enteredToken,
			Email:    enteredEmail,
			Password: enteredPassword,
		},
		Admins:          []string{enteredAdmin},
		CommandPrefix:   cdCommandPrefix,
		AllowSkipping:   cdAllowSkipping,
		ScanOwnMessages: cdScanOwnMessages,

		PresenceEnabled: cdPresenceEnabled,
		PresenceStatus:  cdPresenceStatus,
		PresenceType:    cdPresenceType,

		GithubUpdateChecking: cdGithubUpdateChecking,
		DebugOutput:          cdDebugOutput,
	}

	// Import old config
	if _, err := os.Stat("config.ini"); err == nil {
		log.Println(logPrefixSetup, color.HiGreenString("Detected config.ini from Seklfreak's discord-image-downloader-go, importing..."))
		cfg, err := ini.Load("config.ini")
		if err != nil {
			log.Println(logPrefixSetup, color.HiRedString("Unable to read your old config file:\t%s", err))
			cfg = ini.Empty()
		} else {
			// Import old ini

			// Auth
			if cfg.Section("auth").HasKey("token") {
				defaultConfig.Credentials.Token = cfg.Section("auth").Key("token").String()
				log.Println(color.GreenString("IMPORTED token:\t\t\t%s", defaultConfig.Credentials.Token))
			} else {
				defaultConfig.Credentials.Token = ""
			}
			if cfg.Section("auth").HasKey("email") {
				defaultConfig.Credentials.Email = cfg.Section("auth").Key("email").String()
				log.Println(color.GreenString("IMPORTED email:\t\t\t%s", defaultConfig.Credentials.Email))
			} else {
				defaultConfig.Credentials.Email = ""
			}
			if cfg.Section("auth").HasKey("password") {
				defaultConfig.Credentials.Password = cfg.Section("auth").Key("password").String()
				log.Println(color.GreenString("IMPORTED password:\t\t\t%s", defaultConfig.Credentials.Password))
			} else {
				defaultConfig.Credentials.Password = ""
			}
			if cfg.Section("google").HasKey("client credentials json") {
				defaultConfig.Credentials.GoogleDriveCredentialsJSON = cfg.Section("google").Key("client credentials json").String()
				log.Println(color.GreenString("IMPORTED Google Drive Credentials:\t\t\t%s", defaultConfig.Credentials.GoogleDriveCredentialsJSON))
			}
			if cfg.Section("flickr").HasKey("api key") {
				defaultConfig.Credentials.FlickrApiKey = cfg.Section("flickr").Key("api key").String()
				log.Println(color.GreenString("IMPORTED Flickr API Key:\t\t\t%s", defaultConfig.Credentials.FlickrApiKey))
			}
			if cfg.Section("twitter").HasKey("consumer key") {
				defaultConfig.Credentials.TwitterConsumerKey = cfg.Section("twitter").Key("consumer key").String()
				log.Println(color.GreenString("IMPORTED Twitter Consumer Key:\t\t\t%s", defaultConfig.Credentials.TwitterConsumerKey))
			}
			if cfg.Section("twitter").HasKey("consumer secret") {
				defaultConfig.Credentials.TwitterConsumerSecret = cfg.Section("twitter").Key("consumer secret").String()
				log.Println(color.GreenString("IMPORTED Twitter Consumer Secret:\t\t\t%s", defaultConfig.Credentials.TwitterConsumerSecret))
			}
			if cfg.Section("twitter").HasKey("access token") {
				defaultConfig.Credentials.TwitterAccessToken = cfg.Section("twitter").Key("access token").String()
				log.Println(color.GreenString("IMPORTED Twitter Access Token:\t\t\t%s", defaultConfig.Credentials.TwitterAccessToken))
			}
			if cfg.Section("twitter").HasKey("access secret") {
				defaultConfig.Credentials.TwitterAccessTokenSecret = cfg.Section("twitter").Key("access secret").String()
				log.Println(color.GreenString("IMPORTED Twitter Access Token Secret:\t\t\t%s", defaultConfig.Credentials.TwitterAccessTokenSecret))
			}

			// General
			if cfg.Section("general").HasKey("max download retries") {
				defaultConfig.DownloadRetryMax = cfg.Section("general").Key("max download retries").MustInt()
				log.Println(color.GreenString("IMPORTED Max Download Retries:\t%d", defaultConfig.DownloadRetryMax))
			}
			if cfg.Section("general").HasKey("download timeout") {
				defaultConfig.DownloadTimeout = cfg.Section("general").Key("download timeout").MustInt()
				log.Println(color.GreenString("IMPORTED Download Timeout:\t\t%d", defaultConfig.DownloadRetryMax))
			}

			// Status
			if cfg.Section("status").HasKey("status enabled") {
				defaultConfig.PresenceEnabled = cfg.Section("status").Key("status enabled").MustBool()
				log.Println(color.GreenString("IMPORTED Presence Enabled:\t\t%s", boolS(defaultConfig.PresenceEnabled)))
			}
			if cfg.Section("status").HasKey("status type") {
				defaultConfig.PresenceStatus = cfg.Section("status").Key("status type").String()
				log.Println(color.GreenString("IMPORTED Presence Status:\t\t%s", defaultConfig.PresenceStatus))
			}
			if cfg.Section("status").HasKey("status label") {
				defaultConfig.PresenceType = discordgo.GameType(cfg.Section("status").Key("status label").MustInt())
				log.Println(color.GreenString("IMPORTED Presence Type:\t\t%d", defaultConfig.PresenceType))
			}

			// Channels
			InteractiveChannelWhitelist := cfg.Section("interactive channels").KeysHash()
			for key := range InteractiveChannelWhitelist {
				newChannel := configurationAdminChannel{
					ChannelID: key,
				}
				log.Println(color.GreenString("IMPORTED Admin Channel:\t\t%s", key))
				defaultConfig.AdminChannels = append(defaultConfig.AdminChannels, newChannel)
			}
			ChannelWhitelist := cfg.Section("channels").KeysHash()
			for key, value := range ChannelWhitelist {
				newChannel := configurationChannel{
					ChannelID:   key,
					Destination: value,
				}
				log.Println(color.GreenString("IMPORTED Channel:\t\t\t%s to \"%s\"", key, value))
				defaultConfig.Channels = append(defaultConfig.Channels, newChannel)
			}
		}
		log.Println(logPrefixSetup, color.HiGreenString("Finished importing config.ini from Seklfreak's discord-image-downloader-go!"))
	} else {
		baseChannel := configurationChannel{
			ChannelID:   enteredBaseChannel,
			Destination: enteredBaseDestination,

			Enabled:       &ccdEnabled,
			AllowCommands: &ccdAllowCommands,
			ErrorMessages: &ccdErrorMessages,
			ScanEdits:     &ccdScanEdits,
			IgnoreBots:    &ccdIgnoreBots,

			UpdatePresence:      &ccdUpdatePresence,
			ReactWhenDownloaded: &ccdReactWhenDownloaded,

			DivideFoldersByType: &ccdDivideFoldersByType,
			SaveImages:          &ccdSaveImages,
			SaveVideos:          &ccdSaveVideos,
		}
		defaultConfig.Channels = append(defaultConfig.Channels, baseChannel)

		baseAdminChannel := configurationAdminChannel{
			ChannelID: "REPLACE_WITH_DISCORD_CHANNEL_ID_FOR_ADMIN_COMMANDS",
		}
		defaultConfig.AdminChannels = append(defaultConfig.AdminChannels, baseAdminChannel)

		//TODO: Improve, this is very crude, I just wanted *something* for this.
		log.Print(color.HiCyanString("Would you like to enter settings info now? [Y/N]: "))
		reader := bufio.NewReader(os.Stdin)
		inputCredsYN, _ := reader.ReadString('\n')
		inputCredsYN = strings.ReplaceAll(inputCredsYN, "\n", "")
		inputCredsYN = strings.ReplaceAll(inputCredsYN, "\r", "")
		if strings.Contains(strings.ToLower(inputCredsYN), "y") {
		EnterCreds:
			log.Print(color.HiCyanString("Token or Login? [\"token\"/\"login\"]: "))
			inputCreds, _ := reader.ReadString('\n')
			inputCreds = strings.ReplaceAll(inputCreds, "\n", "")
			inputCreds = strings.ReplaceAll(inputCreds, "\r", "")
			if strings.Contains(strings.ToLower(inputCreds), "token") {
			EnterToken:
				log.Print(color.HiCyanString("Enter token: "))
				inputToken, _ := reader.ReadString('\n')
				inputToken = strings.ReplaceAll(inputToken, "\n", "")
				inputToken = strings.ReplaceAll(inputToken, "\r", "")
				if inputToken != "" {
					enteredToken = inputToken
				} else {
					log.Println(color.HiRedString("Please input token..."))
					goto EnterToken
				}
			} else if strings.Contains(strings.ToLower(inputCreds), "login") {
			EnterEmail:
				log.Print(color.HiCyanString("Enter email: "))
				inputEmail, _ := reader.ReadString('\n')
				inputEmail = strings.ReplaceAll(inputEmail, "\n", "")
				inputEmail = strings.ReplaceAll(inputEmail, "\r", "")
				if strings.Contains(inputEmail, "@") {
					enteredEmail = inputEmail
				EnterPassword:
					log.Print(color.HiCyanString("Enter password: "))
					inputPassword, _ := reader.ReadString('\n')
					inputPassword = strings.ReplaceAll(inputPassword, "\n", "")
					inputPassword = strings.ReplaceAll(inputPassword, "\r", "")
					if inputPassword != "" {
						enteredPassword = inputPassword
					} else {
						log.Println(color.HiRedString("Please input password..."))
						goto EnterPassword
					}
				} else {
					log.Println(color.HiRedString("Please input email..."))
					goto EnterEmail
				}
			} else {
				log.Println(color.HiRedString("Please input \"token\" or \"login\"..."))
				goto EnterCreds
			}

		EnterAdmin:
			log.Print(color.HiCyanString("Input your Discord User ID: "))
			inputAdmin, _ := reader.ReadString('\n')
			inputAdmin = strings.ReplaceAll(inputAdmin, "\n", "")
			inputAdmin = strings.ReplaceAll(inputAdmin, "\r", "")
			if isNumeric(inputAdmin) {
				enteredAdmin = inputAdmin
			} else {
				log.Println(color.HiRedString("Please input your Discord User ID..."))
				goto EnterAdmin
			}

			//TODO: Base channel setup? Would be kind of annoying and may limit options
			//TODO: Admin channel setup?
		}
	}

	log.Println(logPrefixHelper, color.MagentaString("The default settings will be missing some options to avoid clutter."))
	log.Println(logPrefixHelper, color.HiMagentaString("There are MANY MORE SETTINGS! If you would like to maximize customization, see the GitHub README for all available settings."))

	defaultJSON, err := json.MarshalIndent(defaultConfig, "", "\t")
	if err != nil {
		log.Println(logPrefixSetup, color.HiRedString("Failed to format new settings...\t%s", err))
	} else {
		err := ioutil.WriteFile(configPath, defaultJSON, 0644)
		if err != nil {
			log.Println(logPrefixSetup, color.HiRedString("Failed to save new settings file...\t%s", err))
		} else {
			log.Println(logPrefixSetup, color.HiYellowString("Created new settings file..."))
			log.Println(logPrefixSetup, color.HiYellowString("Please save your credentials & info into \"%s\" then restart...", configPath))
			log.Println(logPrefixHelper, color.MagentaString("You DO NOT NEED `Token` *AND* `Email`+`Password`, just one OR the other."))
			log.Println(logPrefixHelper, color.MagentaString("See README on GitHub for help and more info..."))
		}
	}
}

func channelDefault(channel *configurationChannel) {
	// These have to use the default variables since literal values and consts can't be set to the pointers

	// Setup
	if channel.Enabled == nil {
		channel.Enabled = &ccdEnabled
	}
	if channel.AllowCommands == nil {
		channel.AllowCommands = &ccdAllowCommands
	}
	if channel.ErrorMessages == nil {
		channel.ErrorMessages = &ccdErrorMessages
	}
	if channel.ScanEdits == nil {
		channel.ScanEdits = &ccdScanEdits
	}
	if channel.IgnoreBots == nil {
		channel.IgnoreBots = &ccdIgnoreBots
	}
	// Appearance
	if channel.UpdatePresence == nil {
		channel.UpdatePresence = &ccdUpdatePresence
	}
	if channel.ReactWhenDownloaded == nil {
		channel.ReactWhenDownloaded = &ccdReactWhenDownloaded
	}
	if channel.ReactWhenDownloadedEmoji == nil {
		channel.ReactWhenDownloadedEmoji = &ccdReactWhenDownloadedEmoji
	}
	if channel.BlacklistReactEmojis == nil {
		channel.BlacklistReactEmojis = &ccdBlacklistReactEmojis
	}
	if channel.TypeWhileProcessing == nil {
		channel.TypeWhileProcessing = &ccdTypeWhileProcessing
	}
	// Rules for Access
	if channel.UsersAllWhitelisted == nil {
		channel.UsersAllWhitelisted = &ccdUsersAllWhitelisted
	}
	// Rules for Saving
	if channel.DivideFoldersByServer == nil {
		channel.DivideFoldersByServer = &ccdDivideFoldersByServer
	}
	if channel.DivideFoldersByChannel == nil {
		channel.DivideFoldersByChannel = &ccdDivideFoldersByChannel
	}
	if channel.DivideFoldersByUser == nil {
		channel.DivideFoldersByUser = &ccdDivideFoldersByUser
	}
	if channel.DivideFoldersByType == nil {
		channel.DivideFoldersByType = &ccdDivideFoldersByType
	}
	if channel.SaveImages == nil {
		channel.SaveImages = &ccdSaveImages
	}
	if channel.SaveVideos == nil {
		channel.SaveVideos = &ccdSaveVideos
	}
	if channel.SaveAudioFiles == nil {
		channel.SaveAudioFiles = &ccdSaveAudioFiles
	}
	if channel.SaveTextFiles == nil {
		channel.SaveTextFiles = &ccdSaveTextFiles
	}
	if channel.SaveOtherFiles == nil {
		channel.SaveOtherFiles = &ccdSaveOtherFiles
	}
	if channel.SavePossibleDuplicates == nil {
		channel.SavePossibleDuplicates = &ccdSavePossibleDuplicates
	}
	if channel.ExtensionBlacklist == nil {
		channel.ExtensionBlacklist = &ccdExtensionBlacklist
	}
}

//#region Channel Checks/Returns

func isChannelRegistered(ChannelID string) bool {
	for _, item := range config.Channels {
		// Single Channel Config
		if ChannelID == item.ChannelID {
			return true
		}
		// Multi-Channel Config
		if item.ChannelIDs != nil {
			for _, subchannel := range *item.ChannelIDs {
				if ChannelID == subchannel {
					return true
				}
			}
		}
	}
	if config.AllChannels != nil {
		if config.AllChannelsBlacklist != nil {
			if stringInSlice(ChannelID, *config.AllChannelsBlacklist) {
				return false
			}
		}
		return true
	}
	return false
}

func getChannelConfig(ChannelID string) configurationChannel {
	for _, item := range config.Channels {
		// Single Channel Config
		if ChannelID == item.ChannelID {
			return item
		}
		// Multi-Channel Config
		if item.ChannelIDs != nil {
			for _, subchannel := range *item.ChannelIDs {
				if ChannelID == subchannel {
					return item
				}
			}
		}
	}
	if config.AllChannels != nil {
		return *config.AllChannels
	}
	return configurationChannel{}
}

func isAdminChannelRegistered(ChannelID string) bool {
	if config.AdminChannels != nil {
		for _, item := range config.AdminChannels {
			if ChannelID == item.ChannelID {
				return true
			}
		}
	}
	return false
}

func getAdminChannelConfig(ChannelID string) configurationAdminChannel {
	if config.AdminChannels != nil {
		for _, item := range config.AdminChannels {
			if ChannelID == item.ChannelID {
				return item
			}
		}
	}
	return configurationAdminChannel{}
}

func isCommandableChannel(m *discordgo.Message) bool {
	if isAdminChannelRegistered(m.ChannelID) {
		return true
	} else if isChannelRegistered(m.ChannelID) {
		channelConfig := getChannelConfig(m.ChannelID)
		if *channelConfig.AllowCommands || isBotAdmin(m) {
			return true
		}
	}
	return false
}

func isGlobalCommandAllowed(m *discordgo.Message) bool {
	if config.AllowGlobalCommands || isCommandableChannel(m) {
		return true
	}
	return false
}

func getBoundChannels() []string {
	var channels []string
	for _, item := range config.Channels {
		if item.ChannelID != "" {
			if !stringInSlice(item.ChannelID, channels) {
				channels = append(channels, item.ChannelID)
			}
		} else if *item.ChannelIDs != nil {
			for _, subchannel := range *item.ChannelIDs {
				if !stringInSlice(subchannel, channels) {
					channels = append(channels, subchannel)
				}
			}
		}
	}
	return channels
}

func getBoundChannelsCount() int {
	return len(getBoundChannels())
}

//#endregion
