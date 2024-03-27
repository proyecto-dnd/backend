package report

import (
	"bytes"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/proyecto-dnd/backend/internal/attackEvent"
	"github.com/proyecto-dnd/backend/internal/dice_event"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	tradeevent "github.com/proyecto-dnd/backend/internal/tradeEvent"
)

type ReportGenerator struct {
	tradeEventService  tradeevent.ServiceTradeEvent
	attackEventService attackEvent.AttackEventService
	diceEventService   dice_event.DiceEventService
}

func NewReportGenerator(tradeEventService tradeevent.ServiceTradeEvent, attackEventService attackEvent.AttackEventService, diceEventService dice_event.DiceEventService) *ReportGenerator {
	return &ReportGenerator{
		tradeEventService:  tradeEventService,
		attackEventService: attackEventService,
		diceEventService:   diceEventService,
	}
}

func (r *ReportGenerator) GenerateSessionReport(id int) (bytes.Buffer, error) {
	excelFile := excelize.NewFile()
	tradeSheetIndex := excelFile.NewSheet("Trade Events")
	characterTradeSheetIndex := excelFile.NewSheet("Character X Trades")
	attackEventSheetIndex := excelFile.NewSheet("Attack Event")
	affectedByAttackEventSheetIndex := excelFile.NewSheet("Affected By Attack Event")
	
	tradeEvents, err := r.tradeEventService.GetBySessionId(id)
	if err != nil {
		return bytes.Buffer{}, err
	}
	attackEvents, err := r.attackEventService.GetEventsBySessionId(id)
	if err != nil {
		return bytes.Buffer{}, err
	}
	// Todo add dice events

	generateCharacterTradeHeaders(excelFile)
	generateTradeEventHeaders(excelFile)
	for index1, tradeEvent := range tradeEvents {
		for index2, characterTrade := range tradeEvent.TradingItems {
			insertCharacterTradeRow(excelFile, characterTrade, index2)
			excelFile.SetActiveSheet(characterTradeSheetIndex)

		}
		insertTradeEventRow(excelFile, tradeEvent, index1)
		excelFile.SetActiveSheet(tradeSheetIndex)
	}

	generateAttackEventHeaders(excelFile)
	generateAffectedByAttackEventHeaders(excelFile)
	for index1, attackEvent := range attackEvents {
		for index2, affectedByAttackEvent := range attackEvent.Affected {
			insertAffectedByAttackEventRow(excelFile, affectedByAttackEvent, index2)
			excelFile.SetActiveSheet(affectedByAttackEventSheetIndex)
		}
		insertAttackEventRow(excelFile, attackEvent, index1)
		excelFile.SetActiveSheet(attackEventSheetIndex)
	}

	
	if err := excelFile.SaveAs("report.xlsx"); err != nil {
		return bytes.Buffer{}, err
	}
	return bytes.Buffer{}, nil
}

func generateCharacterTradeHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Trades", "A1", "character_trade_id")
	excelFile.SetCellValue("Character X Trades", "B1", "trade_event_id")
	excelFile.SetCellValue("Character X Trades", "C1", "weapon_x_character")
	excelFile.SetCellValue("Character X Trades", "D1", "item_x_character")
	excelFile.SetCellValue("Character X Trades", "E1", "armor_x_character")
	excelFile.SetCellValue("Character X Trades", "F1", "item_owner")
	excelFile.SetCellValue("Character X Trades", "G1", "item_reciever")
	excelFile.SetCellValue("Character X Trades", "H1", "quantity")
	excelFile.SetCellValue("Character X Trades", "I1", "item_name")
	excelFile.SetCellValue("Character X Trades", "J1", "item_type")
}

func insertCharacterTradeRow(excelFile *excelize.File, characterTrade domain.CharacterTrade, i int) {
	excelFile.SetCellValue("Character X Trades", "A"+strconv.Itoa(i+1), characterTrade.CharacterTrade_Id)
	excelFile.SetCellValue("Character X Trades", "B"+strconv.Itoa(i+1), characterTrade.TradeEvent_Id)
	excelFile.SetCellValue("Character X Trades", "C"+strconv.Itoa(i+1), characterTrade.WeaponXCharacter)
	excelFile.SetCellValue("Character X Trades", "D"+strconv.Itoa(i+1), characterTrade.ItemXCharacter)
	excelFile.SetCellValue("Character X Trades", "E"+strconv.Itoa(i+1), characterTrade.ArmorXCharacter)
	excelFile.SetCellValue("Character X Trades", "F"+strconv.Itoa(i+1), characterTrade.ItemOwner)
	excelFile.SetCellValue("Character X Trades", "G"+strconv.Itoa(i+1), characterTrade.ItemReciever)
	excelFile.SetCellValue("Character X Trades", "H"+strconv.Itoa(i+1), characterTrade.Quantity)
	excelFile.SetCellValue("Character X Trades", "I"+strconv.Itoa(i+1), characterTrade.ItemName)
	excelFile.SetCellValue("Character X Trades", "J"+strconv.Itoa(i+1), characterTrade.ItemType)
}

func generateTradeEventHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Trade Event", "A1", "trade_event_id")
	excelFile.SetCellValue("Trade Event", "B1", "session_id")
	excelFile.SetCellValue("Trade Event", "C1", "sender")
	excelFile.SetCellValue("Trade Event", "D1", "receiver")
	excelFile.SetCellValue("Trade Event", "E1", "description")
	excelFile.SetCellValue("Trade Event", "F1", "timestamp")
}

func insertTradeEventRow(excelFile *excelize.File, tradeEvent domain.TradeEvent, i int) {
	excelFile.SetCellValue("Trade Event", "A"+strconv.Itoa(i+1), tradeEvent.TradeEvent_Id)
	excelFile.SetCellValue("Trade Event", "B"+strconv.Itoa(i+1), tradeEvent.Session_Id)
	excelFile.SetCellValue("Trade Event", "C"+strconv.Itoa(i+1), tradeEvent.Sender)
	excelFile.SetCellValue("Trade Event", "D"+strconv.Itoa(i+1), tradeEvent.Receiver)
	excelFile.SetCellValue("Trade Event", "E"+strconv.Itoa(i+1), tradeEvent.Description)
	excelFile.SetCellValue("Trade Event", "F"+strconv.Itoa(i+1), tradeEvent.Timestamp)
}

func generateAttackEventHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Attack Event", "A1", "attack_event_id")
	excelFile.SetCellValue("Attack Event", "B1", "type")
	excelFile.SetCellValue("Attack Event", "C1", "environment")
	excelFile.SetCellValue("Attack Event", "D1", "session_id")
	excelFile.SetCellValue("Attack Event", "E1", "event_protagonist_id")
	excelFile.SetCellValue("Attack Event", "F1", "user_id")
	excelFile.SetCellValue("Attack Event", "G1", "campaign_id")
	excelFile.SetCellValue("Attack Event", "H1", "image_url")
	excelFile.SetCellValue("Attack Event", "I1", "name")
	excelFile.SetCellValue("Attack Event", "J1", "race")
	excelFile.SetCellValue("Attack Event", "K1", "class")
	excelFile.SetCellValue("Attack Event", "L1", "level")
	excelFile.SetCellValue("Attack Event", "M1", "hit_points")
	excelFile.SetCellValue("Attack Event", "N1", "event_resolution")
	excelFile.SetCellValue("Attack Event", "O1", "weapon_id")
	excelFile.SetCellValue("Attack Event", "P1", "spell_id")
	excelFile.SetCellValue("Attack Event", "Q1", "dmg_type")
	excelFile.SetCellValue("Attack Event", "R1", "description")
	excelFile.SetCellValue("Attack Event", "S1", "timestamp")
}

func insertAttackEventRow(excelFile *excelize.File, attackEvent dto.ResponseEventDto, i int) {
	excelFile.SetCellValue("Attack Event", "A"+strconv.Itoa(i+1), attackEvent.AttackEventId)
	excelFile.SetCellValue("Attack Event", "B"+strconv.Itoa(i+1), attackEvent.Type)
	excelFile.SetCellValue("Attack Event", "C"+strconv.Itoa(i+1), attackEvent.Environment)
	excelFile.SetCellValue("Attack Event", "D"+strconv.Itoa(i+1), attackEvent.Session.SessionId)
	excelFile.SetCellValue("Attack Event", "E"+strconv.Itoa(i+1), attackEvent.EventProtagonist.CharacterId)
	excelFile.SetCellValue("Attack Event", "F"+strconv.Itoa(i+1), attackEvent.EventProtagonist.UserId)
	excelFile.SetCellValue("Attack Event", "G"+strconv.Itoa(i+1), attackEvent.EventProtagonist.CampaignID)
	excelFile.SetCellValue("Attack Event", "H"+strconv.Itoa(i+1), attackEvent.EventProtagonist.ImageUrl)
	excelFile.SetCellValue("Attack Event", "I"+strconv.Itoa(i+1), attackEvent.EventProtagonist.Name)
	excelFile.SetCellValue("Attack Event", "J"+strconv.Itoa(i+1), attackEvent.EventProtagonist.Race)
	excelFile.SetCellValue("Attack Event", "K"+strconv.Itoa(i+1), attackEvent.EventProtagonist.Class)
	excelFile.SetCellValue("Attack Event", "L"+strconv.Itoa(i+1), attackEvent.EventProtagonist.Level)
	excelFile.SetCellValue("Attack Event", "M"+strconv.Itoa(i+1), attackEvent.EventProtagonist.HitPoints)
	excelFile.SetCellValue("Attack Event", "N"+strconv.Itoa(i+1), attackEvent.EventResolution)
	excelFile.SetCellValue("Attack Event", "O"+strconv.Itoa(i+1), attackEvent.Weapon)
	excelFile.SetCellValue("Attack Event", "P"+strconv.Itoa(i+1), attackEvent.Spell)
	excelFile.SetCellValue("Attack Event", "Q"+strconv.Itoa(i+1), attackEvent.DmgType)
	excelFile.SetCellValue("Attack Event", "R"+strconv.Itoa(i+1), attackEvent.Description)
	excelFile.SetCellValue("Attack Event", "S"+strconv.Itoa(i+1), attackEvent.TimeStamp)
}

func generateAffectedByAttackEventHeaders(excelFile *excelize.File){
	excelFile.SetCellValue("Affected By Attack Event", "A1", "character_id")
	excelFile.SetCellValue("Affected By Attack Event", "B1", "user_id")
	excelFile.SetCellValue("Affected By Attack Event", "C1", "campaign_id")
	excelFile.SetCellValue("Affected By Attack Event", "D1", "image_url")
	excelFile.SetCellValue("Affected By Attack Event", "E1", "name")
	excelFile.SetCellValue("Affected By Attack Event", "F1", "race")
	excelFile.SetCellValue("Affected By Attack Event", "G1", "class")
	excelFile.SetCellValue("Affected By Attack Event", "H1", "level")
	excelFile.SetCellValue("Affected By Attack Event", "I1", "hit_points")
}

func insertAffectedByAttackEventRow(excelFile *excelize.File, affectedByAttackEvent dto.CharacterCardDto, i int) {
	excelFile.SetCellValue("Affected By Attack Event", "A"+strconv.Itoa(i+1), affectedByAttackEvent.CharacterId)
	excelFile.SetCellValue("Affected By Attack Event", "B"+strconv.Itoa(i+1), affectedByAttackEvent.UserId)
	excelFile.SetCellValue("Affected By Attack Event", "C"+strconv.Itoa(i+1), affectedByAttackEvent.CampaignID)
	excelFile.SetCellValue("Affected By Attack Event", "D"+strconv.Itoa(i+1), affectedByAttackEvent.ImageUrl)
	excelFile.SetCellValue("Affected By Attack Event", "E"+strconv.Itoa(i+1), affectedByAttackEvent.Name)
	excelFile.SetCellValue("Affected By Attack Event", "F"+strconv.Itoa(i+1), affectedByAttackEvent.Race)
	excelFile.SetCellValue("Affected By Attack Event", "G"+strconv.Itoa(i+1), affectedByAttackEvent.Class)
	excelFile.SetCellValue("Affected By Attack Event", "H"+strconv.Itoa(i+1), affectedByAttackEvent.Level)
	excelFile.SetCellValue("Affected By Attack Event", "I"+strconv.Itoa(i+1), affectedByAttackEvent.HitPoints)
}