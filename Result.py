class Result(object):

	def __init__(self,
	             string_date=None,  # TODO replace string date with datetime object
	             balls=list(),
	             lucky_ball=None,
	             number_winner=0,
	             winner_prize=0,
	             next_lottery_string_date=None,  # TODO replace string date with datetime object
	             next_lottery_prize=0):
		self.string_date = string_date
		self.ordered_balls = balls
		self.lucky_ball = lucky_ball
		self.number_winner = number_winner
		self.winner_prize = winner_prize
		self.next_lottery_date = next_lottery_string_date
		self.next_lottery_prize = next_lottery_prize

	def is_winning(self, list_winning_balls, winning_lucky_ball):
		sorted_winning_balls = sorted(list_winning_balls, key=int)
		return sorted_winning_balls == self.ordered_balls and winning_lucky_ball == self.lucky_ball

	def __str__(self):
		return "Résultats du %s \n" \
		       "Numéros : %s \n" \
		       "Numéro chance : %s \n" \
		       "%s \n" \
		       "Le prochain tirage sera le %s pour un montant de %i €.\n" % \
		       (self.string_date, \
		        " ".join(self.ordered_balls), \
		        self.lucky_ball, \
		        Result.__get_current_winner_string__(), \
		        self.next_lottery_date, \
		        self.next_lottery_prize)

	def __get_current_winner_string__(self):
		current_winner = "Le jackpot n'a pas été remporté lors de ce tirage !"
		if self.number_winner == 1:
			current_winner = "Un joueur a remporté le jackpot d'un montant de %i €" % self.winner_prize
		elif self.number_winner > 1:
			current_winner = "Le jackpot a été remporté par %i joueurs, ils se partagent %i €" % (
				self.number_winner, self.winner_prize)
		return current_winner
