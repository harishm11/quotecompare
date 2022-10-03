package rtgengine

import (
	"time"

	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/tables/ratingtables"
)

func GetRatingFactor(step RateStep) (float32, time.Time) {
	var factorRow ratingtables.RateFactors

	tempcvg := step.CvgCodetoGetFctr
	if step.RateVar1Code == "TermLength" {
		step.CvgCodetoGetFctr = "All"
	}
	db := database.DBConn

	db.Table("rate_factors").Where("coverage_code = ? AND rating_item_code = ? AND (rate_var1_code= ? OR rate_var1_code is NULL ) AND	(rate_var1_value =? OR rate_var1_value is NULL ) AND (rate_var2_code= ? OR rate_var2_code is NULL )AND (rate_var2_value =? OR rate_var2_value is NULL )AND (rate_var3_code= ? OR rate_var3_code is NULL ) AND (rate_var3_value =? OR rate_var3_value is NULL )AND (rate_var4_code= ? OR rate_var4_code is NULL )AND (rate_var4_value =? OR rate_var4_value is NULL )AND(rate_var5_code= ? OR rate_var5_code is NULL ) AND	(rate_var5_value =? OR rate_var5_value is NULL ) AND (rate_var6_code= ? OR rate_var6_code is NULL )AND (rate_var6_value =? OR rate_var6_value is NULL )AND (rate_var7_code= ? OR rate_var7_code is NULL ) AND (rate_var7_value =? OR rate_var7_value is NULL )AND (rate_var8_code= ? OR rate_var8_code is NULL )AND (rate_var8_value =? OR rate_var8_value is NULL ) AND rate_new_business_eff_date <= ? AND rate_activation_date <= ? AND retired_row = ? ",
		step.CvgCodetoGetFctr, step.RatingItemCode,
		step.RateVar1Code, step.RateVar1Value,
		step.RateVar2Code, step.RateVar2Value,
		step.RateVar3Code, step.RateVar3Value,
		step.RateVar4Code, step.RateVar4Value,
		step.RateVar5Code, step.RateVar5Value,
		step.RateVar6Code, step.RateVar6Value,
		step.RateVar7Code, step.RateVar7Value,
		step.RateVar8Code, step.RateVar8Value, step.RateEffDate, step.RatebookActivationDate, false).Select("rate_factor,rate_activation_date").Order("rate_activation_date desc").Last(&factorRow)
	step.RateFactor = factorRow.RateFactor
	step.RateActivationDate = factorRow.RateActivationDate

	step.CvgCodetoGetFctr = tempcvg
	if step.RateFactor == 0 {
		step.RateFactor = 1.0
		step.RateActivationDate = time.Now()
	}
	return step.RateFactor, step.RateActivationDate

}
