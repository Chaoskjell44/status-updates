package sharder
    
import (  
   "github.com/ticketsbot-cloud/gdl/cache"   
   "github.com/ticketsbot-cloud/gdl/gateway"  
   "github.com/ticketsbot-cloud/gdl/gateway/payloads/events"  
   "github.com/ticketsbot-cloud/gdl/objects/user"  
   "github.com/TicketsBot-cloud/status-updates/internal/config"
)    
    
func main() { 
    shardOptions := gateway.ShardOptions{ 
        ShardCount: gateway.ShardCount{  
            Total:   1,  
            Lowest:  0,  // Inclusive
            Highest: 1,  // Exclusive
        },
        RateLimitStore: ratelimit.NewMemoryStore(), // ratelimit.NewRedisStore() is also available
        
        // We can choose exactly what we want to cache. Remember the zero value of a bool is false!
        CacheFactory: cache.MemoryCacheFactory(cache.CacheOptions{
            Guilds:      false,
            Users:       false,
            Members:     false,
            Channels:    false,
            Roles:       false,
            Emojis:      false,
            VoiceStates: false,
        }),
        GuildSubscriptions: false,
        Presence: user.BuildStatus(user.ActivityTypePlaying, "DM for help | t!help"), // Set the status of the bot
    }

    token := "d.config.Discord.Token"  
    sm := gateway.NewShardManager(token, shardOptions)
    sm.RegisterListeners(echoListener)  
    sm.Connect()  
    sm.WaitForInterrupt()  
}    