type RankingUpdateDto struct {
    Ranking *map[string]int
    TargetPlayer string
    Calling string
    NowRanking int
}

func updateRankingMapping(update RankingUpdateDto) {
    (*update.Ranking)[update.TargetPlayer] = update.NowRanking
    (*update.Ranking)[update.Calling] = update.NowRanking - 1
}

func solution(players []string, callings []string) []string {
    ranking := make(map[string]int)
    for idx, players := range players {
        ranking[players] = idx
    }
    for _, calling := range callings {
        nowRanking := ranking[calling]
        if nowRanking != 0 {
            targetPlayer := players[nowRanking - 1]
            players[nowRanking], players[nowRanking - 1] = players[nowRanking - 1], players[nowRanking]
            updateRankingMapping(RankingUpdateDto{
                Ranking: &ranking,
                TargetPlayer: targetPlayer,
                Calling: calling,
                NowRanking: nowRanking,
            })
        }
    }
    return players
}