```kotlin
@MappedSuperclass
abstract class BaseEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long = 0

    @CreationTimestamp
    val createdAt: LocalDateTime = LocalDateTime.MIN

    @UpdateTimestamp
    val updatedAt: LocalDateTime = LocalDateTime.MIN
}

data class Member(
    val id: Long? = null,
    val name: String,
    val email: String?,
    val profileImage: String?,
    val role: MemberRole,
    val socialInfo: SocialInfo,
    val activityScore: Int = 0,
) {
    companion object {
        fun register(
            name: String,
            email: String,
            profileImage: String?,
            role: MemberRole,
            provider: SocialProvider,
            socialId: String,
        ): Member =
            Member(
                name = name,
                email = email,
                profileImage = profileImage,
                role = role,
                socialInfo = SocialInfo(provider, socialId),
            )
    }
    
}


data class SocialInfo(
    val provider: SocialProvider,
    val socialId: String,
)


enum class SocialProvider(
    val description: String,
) {
    GOOGLE("구글"),
    APPLE("애플"),
    KAKAO("카카오"),
}


```