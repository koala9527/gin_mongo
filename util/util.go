package util

type WebcastChatMessage struct {
	Common struct {
		Method        string `json:"method"`
		MsgID         string `json:"msgId"`
		RoomID        string `json:"roomId"`
		IsShowMsg     bool   `json:"isShowMsg"`
		PriorityScore string `json:"priorityScore"`
	} `json:"common"`
	User struct {
		ID          string `json:"id"`
		ShortID     string `json:"shortId"`
		Nickname    string `json:"nickname"`
		Gender      int    `json:"gender"`
		AvatarThumb struct {
			URLList []string `json:"urlList"`
			URI     string   `json:"uri"`
		} `json:"avatarThumb"`
		BadgeImageList []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height,omitempty"`
			Width     string   `json:"width,omitempty"`
			ImageType int      `json:"imageType"`
			Content   struct {
				Name      string `json:"name"`
				FontColor string `json:"fontColor"`
				Level     string `json:"level"`
			} `json:"content,omitempty"`
		} `json:"badgeImageList"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount  string `json:"followerCount"`
			FollowStatus   string `json:"followStatus"`
		} `json:"followInfo"`
		PayGrade struct {
			Level              string `json:"level"`
			NewImIconWithLevel struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newImIconWithLevel"`
			NewLiveIcon struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newLiveIcon"`
		} `json:"payGrade"`
		FansClub struct {
			Data struct {
				ClubName           string `json:"clubName"`
				Level              int    `json:"level"`
				UserFansClubStatus string `json:"userFansClubStatus"`
				Badge              struct {
					Icons struct {
						Num2 struct {
							URLList []string `json:"urlList"`
							URI     string   `json:"uri"`
						} `json:"2"`
					} `json:"icons"`
					Title string `json:"title"`
				} `json:"badge"`
				AnchorID string `json:"anchorId"`
			} `json:"data"`
			PreferData struct {
				Num1 struct {
					ClubName           string `json:"clubName"`
					Level              int    `json:"level"`
					UserFansClubStatus string `json:"userFansClubStatus"`
					Badge              struct {
						Icons struct {
							Num2 struct {
								URLList []string `json:"urlList"`
								URI     string   `json:"uri"`
							} `json:"2"`
						} `json:"icons"`
						Title string `json:"title"`
					} `json:"badge"`
					AnchorID string `json:"anchorId"`
				} `json:"1"`
			} `json:"preferData"`
		} `json:"fansClub"`
		UserAttr struct {
		} `json:"userAttr"`
		DisplayID         string `json:"displayId"`
		SecUID            string `json:"secUid"`
		AuthorizationInfo int    `json:"authorizationInfo"`
		BadgeImageListV2  []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height,omitempty"`
			Width     string   `json:"width,omitempty"`
			ImageType int      `json:"imageType"`
			Content   struct {
				Name      string `json:"name"`
				FontColor string `json:"fontColor"`
				Level     string `json:"level"`
			} `json:"content,omitempty"`
		} `json:"badgeImageListV2"`
		MysteryMan           int    `json:"mysteryMan"`
		DesensitizedNickname string `json:"desensitizedNickname"`
	} `json:"user"`
	Content          string `json:"content"`
	PublicAreaCommon struct {
		UserLabel struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"userLabel"`
		UserConsumeInRoom string `json:"userConsumeInRoom"`
	} `json:"publicAreaCommon"`
}

type WebcastMemberMessage struct {
	Common struct {
		Method      string `json:"method"`
		MsgID       string `json:"msgId"`
		RoomID      string `json:"roomId"`
		IsShowMsg   bool   `json:"isShowMsg"`
		DisplayText struct {
			Key            string `json:"key"`
			DefaultPattern string `json:"defaultPattern"`
			DefaultFormat  struct {
				Color  string `json:"color"`
				Weight int    `json:"weight"`
			} `json:"defaultFormat"`
			Pieces []struct {
				Type   int `json:"type"`
				Format struct {
					Color  string `json:"color"`
					Weight int    `json:"weight"`
				} `json:"format"`
				UserValue struct {
					User struct {
						ID          string `json:"id"`
						ShortID     string `json:"shortId"`
						Nickname    string `json:"nickname"`
						AvatarThumb struct {
							URLList []string `json:"urlList"`
							URI     string   `json:"uri"`
						} `json:"avatarThumb"`
						FollowInfo struct {
							FollowingCount string `json:"followingCount"`
							FollowerCount  string `json:"followerCount"`
						} `json:"followInfo"`
						PayGrade struct {
						} `json:"payGrade"`
						FansClub struct {
							Data struct {
								Badge struct {
									Icons struct {
										Num0 struct {
										} `json:"0"`
									} `json:"icons"`
								} `json:"badge"`
							} `json:"data"`
						} `json:"fansClub"`
						UserAttr struct {
						} `json:"userAttr"`
						DisplayID            string `json:"displayId"`
						SecUID               string `json:"secUid"`
						AuthorizationInfo    int    `json:"authorizationInfo"`
						MysteryMan           int    `json:"mysteryMan"`
						DesensitizedNickname string `json:"desensitizedNickname"`
					} `json:"user"`
				} `json:"userValue"`
			} `json:"pieces"`
		} `json:"displayText"`
		FoldType         string `json:"foldType"`
		AnchorFoldType   string `json:"anchorFoldType"`
		PriorityScore    string `json:"priorityScore"`
		AnchorFoldTypeV2 string `json:"anchorFoldTypeV2"`
	} `json:"common"`
	User struct {
		ID          string `json:"id"`
		ShortID     string `json:"shortId"`
		Nickname    string `json:"nickname"`
		AvatarThumb struct {
			URLList []string `json:"urlList"`
			URI     string   `json:"uri"`
		} `json:"avatarThumb"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount  string `json:"followerCount"`
		} `json:"followInfo"`
		PayGrade struct {
		} `json:"payGrade"`
		FansClub struct {
			Data struct {
				Badge struct {
					Icons struct {
						Num0 struct {
						} `json:"0"`
					} `json:"icons"`
				} `json:"badge"`
			} `json:"data"`
		} `json:"fansClub"`
		UserAttr struct {
		} `json:"userAttr"`
		DisplayID            string `json:"displayId"`
		SecUID               string `json:"secUid"`
		AuthorizationInfo    int    `json:"authorizationInfo"`
		MysteryMan           int    `json:"mysteryMan"`
		DesensitizedNickname string `json:"desensitizedNickname"`
	} `json:"user"`
	MemberCount       string `json:"memberCount"`
	Action            string `json:"action"`
	AnchorDisplayText struct {
		Key            string `json:"key"`
		DefaultPattern string `json:"defaultPattern"`
		DefaultFormat  struct {
			Color  string `json:"color"`
			Weight int    `json:"weight"`
		} `json:"defaultFormat"`
		Pieces []struct {
			Type   int `json:"type"`
			Format struct {
				Color  string `json:"color"`
				Weight int    `json:"weight"`
			} `json:"format"`
			UserValue struct {
				User struct {
					ID          string `json:"id"`
					ShortID     string `json:"shortId"`
					Nickname    string `json:"nickname"`
					AvatarThumb struct {
						URLList []string `json:"urlList"`
						URI     string   `json:"uri"`
					} `json:"avatarThumb"`
					FollowInfo struct {
						FollowingCount string `json:"followingCount"`
						FollowerCount  string `json:"followerCount"`
					} `json:"followInfo"`
					PayGrade struct {
					} `json:"payGrade"`
					FansClub struct {
						Data struct {
							Badge struct {
								Icons struct {
									Num0 struct {
									} `json:"0"`
								} `json:"icons"`
							} `json:"badge"`
						} `json:"data"`
					} `json:"fansClub"`
					UserAttr struct {
					} `json:"userAttr"`
					DisplayID            string `json:"displayId"`
					SecUID               string `json:"secUid"`
					AuthorizationInfo    int    `json:"authorizationInfo"`
					MysteryMan           int    `json:"mysteryMan"`
					DesensitizedNickname string `json:"desensitizedNickname"`
				} `json:"user"`
			} `json:"userValue"`
		} `json:"pieces"`
	} `json:"anchorDisplayText"`
}

type WebcastSocialMessage struct {
	Common struct {
		Method      string `json:"method"`
		MsgID       string `json:"msgId"`
		RoomID      string `json:"roomId"`
		IsShowMsg   bool   `json:"isShowMsg"`
		DisplayText struct {
			Key            string `json:"key"`
			DefaultPattern string `json:"defaultPattern"`
			DefaultFormat  struct {
				Color  string `json:"color"`
				Weight int    `json:"weight"`
			} `json:"defaultFormat"`
			Pieces []struct {
				Type      int `json:"type"`
				UserValue struct {
					User struct {
						ID          string `json:"id"`
						ShortID     string `json:"shortId"`
						Nickname    string `json:"nickname"`
						Gender      int    `json:"gender"`
						Level       int    `json:"level"`
						AvatarThumb struct {
							URLList []string `json:"urlList"`
							URI     string   `json:"uri"`
						} `json:"avatarThumb"`
						BadgeImageList []struct {
							URLList   []string `json:"urlList"`
							URI       string   `json:"uri"`
							Height    string   `json:"height"`
							Width     string   `json:"width"`
							ImageType int      `json:"imageType"`
						} `json:"badgeImageList"`
						FollowInfo struct {
							FollowingCount string `json:"followingCount"`
							FollowerCount  string `json:"followerCount"`
							FollowStatus   string `json:"followStatus"`
						} `json:"followInfo"`
						PayGrade struct {
							Level              string `json:"level"`
							NewImIconWithLevel struct {
								URLList   []string `json:"urlList"`
								URI       string   `json:"uri"`
								Height    string   `json:"height"`
								Width     string   `json:"width"`
								ImageType int      `json:"imageType"`
							} `json:"newImIconWithLevel"`
							NewLiveIcon struct {
								URLList   []string `json:"urlList"`
								URI       string   `json:"uri"`
								Height    string   `json:"height"`
								Width     string   `json:"width"`
								ImageType int      `json:"imageType"`
							} `json:"newLiveIcon"`
						} `json:"payGrade"`
						FansClub struct {
							Data struct {
								Badge struct {
									Icons struct {
										Num0 struct {
										} `json:"0"`
									} `json:"icons"`
								} `json:"badge"`
							} `json:"data"`
						} `json:"fansClub"`
						UserAttr struct {
						} `json:"userAttr"`
						DisplayID         string `json:"displayId"`
						SecUID            string `json:"secUid"`
						AuthorizationInfo int    `json:"authorizationInfo"`
						BadgeImageListV2  []struct {
							URLList   []string `json:"urlList"`
							URI       string   `json:"uri"`
							Height    string   `json:"height"`
							Width     string   `json:"width"`
							ImageType int      `json:"imageType"`
						} `json:"badgeImageListV2"`
						MysteryMan           int    `json:"mysteryMan"`
						DesensitizedNickname string `json:"desensitizedNickname"`
					} `json:"user"`
				} `json:"userValue"`
			} `json:"pieces"`
		} `json:"displayText"`
		PriorityScore string `json:"priorityScore"`
	} `json:"common"`
	User struct {
		ID          string `json:"id"`
		ShortID     string `json:"shortId"`
		Nickname    string `json:"nickname"`
		Gender      int    `json:"gender"`
		Level       int    `json:"level"`
		AvatarThumb struct {
			URLList []string `json:"urlList"`
			URI     string   `json:"uri"`
		} `json:"avatarThumb"`
		BadgeImageList []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height"`
			Width     string   `json:"width"`
			ImageType int      `json:"imageType"`
		} `json:"badgeImageList"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount  string `json:"followerCount"`
			FollowStatus   string `json:"followStatus"`
		} `json:"followInfo"`
		PayGrade struct {
			Level              string `json:"level"`
			NewImIconWithLevel struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newImIconWithLevel"`
			NewLiveIcon struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newLiveIcon"`
		} `json:"payGrade"`
		FansClub struct {
			Data struct {
				Badge struct {
					Icons struct {
						Num0 struct {
						} `json:"0"`
					} `json:"icons"`
				} `json:"badge"`
			} `json:"data"`
		} `json:"fansClub"`
		UserAttr struct {
		} `json:"userAttr"`
		DisplayID         string `json:"displayId"`
		SecUID            string `json:"secUid"`
		AuthorizationInfo int    `json:"authorizationInfo"`
		BadgeImageListV2  []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height"`
			Width     string   `json:"width"`
			ImageType int      `json:"imageType"`
		} `json:"badgeImageListV2"`
		MysteryMan           int    `json:"mysteryMan"`
		DesensitizedNickname string `json:"desensitizedNickname"`
	} `json:"user"`
	Action           string `json:"action"`
	ShareTarget      string `json:"shareTarget"`
	FollowCount      string `json:"followCount"`
	PublicAreaCommon struct {
		UserLabel struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"userLabel"`
	} `json:"publicAreaCommon"`
}

type WebcastLikeMessage struct {
	Common struct {
		Method     string `json:"method"`
		MsgID      string `json:"msgId"`
		RoomID     string `json:"roomId"`
		CreateTime string `json:"createTime"`
		IsShowMsg  bool   `json:"isShowMsg"`
	} `json:"common"`
	Count string `json:"count"`
	Total string `json:"total"`
	User  struct {
		ID          string `json:"id"`
		ShortID     string `json:"shortId"`
		Nickname    string `json:"nickname"`
		Gender      int    `json:"gender"`
		AvatarThumb struct {
			URLList []string `json:"urlList"`
			URI     string   `json:"uri"`
		} `json:"avatarThumb"`
		BadgeImageList []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height"`
			Width     string   `json:"width"`
			ImageType int      `json:"imageType"`
		} `json:"badgeImageList"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount  string `json:"followerCount"`
		} `json:"followInfo"`
		PayGrade struct {
			Level              string `json:"level"`
			NewImIconWithLevel struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newImIconWithLevel"`
			NewLiveIcon struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newLiveIcon"`
		} `json:"payGrade"`
		FansClub struct {
			Data struct {
				Badge struct {
					Icons struct {
						Num0 struct {
						} `json:"0"`
					} `json:"icons"`
				} `json:"badge"`
			} `json:"data"`
		} `json:"fansClub"`
		UserAttr struct {
		} `json:"userAttr"`
		DisplayID         string `json:"displayId"`
		SecUID            string `json:"secUid"`
		AuthorizationInfo int    `json:"authorizationInfo"`
		BadgeImageListV2  []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height"`
			Width     string   `json:"width"`
			ImageType int      `json:"imageType"`
		} `json:"badgeImageListV2"`
	} `json:"user"`
}

type WebcastRoomUserSeqMessage struct {
	Common struct {
		Method     string `json:"method"`
		MsgID      string `json:"msgId"`
		RoomID     string `json:"roomId"`
		CreateTime string `json:"createTime"`
	} `json:"common"`
	Ranks []struct {
		User struct {
			ID          string `json:"id"`
			ShortID     string `json:"shortId"`
			Nickname    string `json:"nickname"`
			AvatarThumb struct {
				URLList []string `json:"urlList"`
				URI     string   `json:"uri"`
			} `json:"avatarThumb"`
			PayGrade struct {
			} `json:"payGrade"`
			AuthorizationInfo int `json:"authorizationInfo"`
		} `json:"user"`
		Rank             string `json:"rank"`
		ScoreDescription string `json:"scoreDescription"`
		ExactlyScore     string `json:"exactlyScore"`
	} `json:"ranks"`
	Total               string `json:"total"`
	TotalUser           string `json:"totalUser"`
	TotalUserStr        string `json:"totalUserStr"`
	TotalStr            string `json:"totalStr"`
	OnlineUserForAnchor string `json:"onlineUserForAnchor"`
	TotalPvForAnchor    string `json:"totalPvForAnchor"`
}

type WebcastGiftMessage struct {
	Common struct {
		Method      string `json:"method"`
		MsgID       string `json:"msgId"`
		RoomID      string `json:"roomId"`
		CreateTime  string `json:"createTime"`
		IsShowMsg   bool   `json:"isShowMsg"`
		Describe    string `json:"describe"`
		DisplayText struct {
			Key            string `json:"key"`
			DefaultPattern string `json:"defaultPattern"`
			DefaultFormat  struct {
				Color  string `json:"color"`
				Weight int    `json:"weight"`
			} `json:"defaultFormat"`
			Pieces []struct {
				Type   int `json:"type"`
				Format struct {
					Color  string `json:"color"`
					Weight int    `json:"weight"`
				} `json:"format,omitempty"`
				UserValue struct {
					User struct {
						ID          string `json:"id"`
						ShortID     string `json:"shortId"`
						Nickname    string `json:"nickname"`
						AvatarThumb struct {
							URLList []string `json:"urlList"`
							URI     string   `json:"uri"`
						} `json:"avatarThumb"`
						BadgeImageList []struct {
							URLList   []string `json:"urlList"`
							URI       string   `json:"uri"`
							Height    string   `json:"height,omitempty"`
							Width     string   `json:"width,omitempty"`
							ImageType int      `json:"imageType"`
							Content   struct {
								Name      string `json:"name"`
								FontColor string `json:"fontColor"`
								Level     string `json:"level"`
							} `json:"content,omitempty"`
						} `json:"badgeImageList"`
						FollowInfo struct {
							FollowingCount string `json:"followingCount"`
							FollowerCount  string `json:"followerCount"`
						} `json:"followInfo"`
						PayGrade struct {
							Level              string `json:"level"`
							NewImIconWithLevel struct {
								URLList   []string `json:"urlList"`
								URI       string   `json:"uri"`
								Height    string   `json:"height"`
								Width     string   `json:"width"`
								ImageType int      `json:"imageType"`
							} `json:"newImIconWithLevel"`
							NewLiveIcon struct {
								URLList   []string `json:"urlList"`
								URI       string   `json:"uri"`
								Height    string   `json:"height"`
								Width     string   `json:"width"`
								ImageType int      `json:"imageType"`
							} `json:"newLiveIcon"`
						} `json:"payGrade"`
						FansClub struct {
							Data struct {
								ClubName           string `json:"clubName"`
								Level              int    `json:"level"`
								UserFansClubStatus string `json:"userFansClubStatus"`
								Badge              struct {
									Icons struct {
										Num2 struct {
											URLList []string `json:"urlList"`
											URI     string   `json:"uri"`
										} `json:"2"`
									} `json:"icons"`
									Title string `json:"title"`
								} `json:"badge"`
								AnchorID string `json:"anchorId"`
							} `json:"data"`
						} `json:"fansClub"`
						UserAttr struct {
						} `json:"userAttr"`
						DisplayID         string `json:"displayId"`
						SecUID            string `json:"secUid"`
						AuthorizationInfo int    `json:"authorizationInfo"`
						BadgeImageListV2  []struct {
							URLList   []string `json:"urlList"`
							URI       string   `json:"uri"`
							Height    string   `json:"height,omitempty"`
							Width     string   `json:"width,omitempty"`
							ImageType int      `json:"imageType"`
							Content   struct {
								Name      string `json:"name"`
								FontColor string `json:"fontColor"`
								Level     string `json:"level"`
							} `json:"content,omitempty"`
						} `json:"badgeImageListV2"`
						MysteryMan           int    `json:"mysteryMan"`
						DesensitizedNickname string `json:"desensitizedNickname"`
					} `json:"user"`
				} `json:"userValue,omitempty"`
				GiftValue struct {
					GiftID  string `json:"giftId"`
					NameRef struct {
						Key            string `json:"key"`
						DefaultPattern string `json:"defaultPattern"`
					} `json:"nameRef"`
				} `json:"giftValue,omitempty"`
				StringValue string `json:"stringValue,omitempty"`
			} `json:"pieces"`
		} `json:"displayText"`
		PriorityScore string `json:"priorityScore"`
	} `json:"common"`
	GiftID      string `json:"giftId"`
	GroupCount  string `json:"groupCount"`
	RepeatCount string `json:"repeatCount"`
	ComboCount  string `json:"comboCount"`
	User        struct {
		ID          string `json:"id"`
		ShortID     string `json:"shortId"`
		Nickname    string `json:"nickname"`
		AvatarThumb struct {
			URLList []string `json:"urlList"`
			URI     string   `json:"uri"`
		} `json:"avatarThumb"`
		BadgeImageList []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height,omitempty"`
			Width     string   `json:"width,omitempty"`
			ImageType int      `json:"imageType"`
			Content   struct {
				Name      string `json:"name"`
				FontColor string `json:"fontColor"`
				Level     string `json:"level"`
			} `json:"content,omitempty"`
		} `json:"badgeImageList"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount  string `json:"followerCount"`
		} `json:"followInfo"`
		PayGrade struct {
			Level              string `json:"level"`
			NewImIconWithLevel struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newImIconWithLevel"`
			NewLiveIcon struct {
				URLList   []string `json:"urlList"`
				URI       string   `json:"uri"`
				Height    string   `json:"height"`
				Width     string   `json:"width"`
				ImageType int      `json:"imageType"`
			} `json:"newLiveIcon"`
		} `json:"payGrade"`
		FansClub struct {
			Data struct {
				ClubName           string `json:"clubName"`
				Level              int    `json:"level"`
				UserFansClubStatus string `json:"userFansClubStatus"`
				Badge              struct {
					Icons struct {
						Num2 struct {
							URLList []string `json:"urlList"`
							URI     string   `json:"uri"`
						} `json:"2"`
					} `json:"icons"`
					Title string `json:"title"`
				} `json:"badge"`
				AnchorID string `json:"anchorId"`
			} `json:"data"`
		} `json:"fansClub"`
		UserAttr struct {
		} `json:"userAttr"`
		DisplayID         string `json:"displayId"`
		SecUID            string `json:"secUid"`
		AuthorizationInfo int    `json:"authorizationInfo"`
		BadgeImageListV2  []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height,omitempty"`
			Width     string   `json:"width,omitempty"`
			ImageType int      `json:"imageType"`
			Content   struct {
				Name      string `json:"name"`
				FontColor string `json:"fontColor"`
				Level     string `json:"level"`
			} `json:"content,omitempty"`
		} `json:"badgeImageListV2"`
		MysteryMan           int    `json:"mysteryMan"`
		DesensitizedNickname string `json:"desensitizedNickname"`
	} `json:"user"`
	RepeatEnd int    `json:"repeatEnd"`
	GroupID   string `json:"groupId"`
	Priority  struct {
		QueueSizes        []string `json:"queueSizes"`
		SelfQueuePriority string   `json:"selfQueuePriority"`
		Priority          string   `json:"priority"`
	} `json:"priority"`
	Gift struct {
		Image struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"image"`
		Describe      string `json:"describe"`
		Duration      string `json:"duration"`
		ID            string `json:"id"`
		ForLinkmic    bool   `json:"forLinkmic"`
		Combo         bool   `json:"combo"`
		Type          int    `json:"type"`
		DiamondCount  int    `json:"diamondCount"`
		GiftLabelIcon struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"giftLabelIcon"`
		Name string `json:"name"`
		Icon struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"icon"`
		GiftScene string `json:"giftScene"`
	} `json:"gift"`
	LogID            string `json:"logId"`
	SendType         string `json:"sendType"`
	PublicAreaCommon struct {
		UserLabel struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"userLabel"`
		UserConsumeInRoom string `json:"userConsumeInRoom"`
	} `json:"publicAreaCommon"`
	TrayDisplayText struct {
		Key            string `json:"key"`
		DefaultPattern string `json:"defaultPattern"`
		DefaultFormat  struct {
			Color         string `json:"color"`
			Weight        int    `json:"weight"`
			UseRemoteClor bool   `json:"useRemoteClor"`
		} `json:"defaultFormat"`
		Pieces []struct {
			Type   int `json:"type"`
			Format struct {
				Color         string `json:"color"`
				Weight        int    `json:"weight"`
				UseRemoteClor bool   `json:"useRemoteClor"`
			} `json:"format"`
			StringValue string `json:"stringValue"`
		} `json:"pieces"`
	} `json:"trayDisplayText"`
	TrayInfo struct {
		TrayBaseImg struct {
			URLList  []string `json:"urlList"`
			URI      string   `json:"uri"`
			AvgColor string   `json:"avgColor"`
		} `json:"trayBaseImg"`
		TrayLevel string `json:"trayLevel"`
	} `json:"trayInfo"`
}

type WebcastRoomIntroMessage struct {
	Common struct {
		Method    string `json:"method"`
		MsgID     string `json:"msgId"`
		IsShowMsg bool   `json:"isShowMsg"`
	} `json:"common"`
	User struct {
		ID          string `json:"id"`
		ShortID     string `json:"shortId"`
		Nickname    string `json:"nickname"`
		Gender      int    `json:"gender"`
		AvatarThumb struct {
			URLList []string `json:"urlList"`
			URI     string   `json:"uri"`
		} `json:"avatarThumb"`
		BadgeImageList []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height"`
			Width     string   `json:"width"`
			ImageType int      `json:"imageType"`
		} `json:"badgeImageList"`
		FollowInfo struct {
			FollowingCount string `json:"followingCount"`
			FollowerCount  string `json:"followerCount"`
		} `json:"followInfo"`
		UserAttr struct {
		} `json:"userAttr"`
		DisplayID         string `json:"displayId"`
		SecUID            string `json:"secUid"`
		AuthorizationInfo int    `json:"authorizationInfo"`
		BadgeImageListV2  []struct {
			URLList   []string `json:"urlList"`
			URI       string   `json:"uri"`
			Height    string   `json:"height"`
			Width     string   `json:"width"`
			ImageType int      `json:"imageType"`
		} `json:"badgeImageListV2"`
	} `json:"user"`
	Intro string   `json:"intro"`
	Label []string `json:"label"`
}
